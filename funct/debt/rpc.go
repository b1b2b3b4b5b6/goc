package debt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/lumosin/goc/tl/cfgt"
	"github.com/lumosin/goc/tl/errt"
)

type response struct {
	ErrorCode int
	Message   string                 `json:"Message,omitempty"`
	Result    map[string]interface{} `json:"Result,omitempty"`
}

type requestArg struct {
	Func   string
	Params interface{}
}

func init() {

	http.HandleFunc("/debug", debFunc)
	http.HandleFunc("/echo", echoFunc)
	cfg := cfgt.New("conf.json")
	port, err := cfg.TakeInt("DebugPort")
	errt.Errpanic(err)
	go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	fmt.Printf("[rpc.go]debug tool work on port[%d]\r\n", port)

}

var funcList = make(map[string]interface{})

func AddFunc(name string, foo interface{}) {
	_, ok := funcList[name]
	if ok {
		panic(fmt.Sprintf("repeat debug func name[%s]", name))
	}
	funcList[name] = foo
	fmt.Printf("[rpc.go]add debug func[%s]\r\n", name)
}

func echoFunc(w http.ResponseWriter, r *http.Request) {
	res := response{ErrorCode: 0, Result: make(map[string]interface{})}
	defer func() {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(jsontool.Encode(&res)))

	}()
	err := r.ParseForm()
	errt.Errpanic(err)
	var m interface{}
	err = json.NewDecoder(r.Body).Decode(&m)
	fmt.Printf("echo[%+v]\r\n", m)
}

func debFunc(w http.ResponseWriter, r *http.Request) {
	res := response{ErrorCode: 0, Result: make(map[string]interface{})}
	defer func() {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(jsontool.Encode(&res)))

	}()

	err := r.ParseForm()
	errt.Errpanic(err)
	var request requestArg
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		res.ErrorCode = -1
		res.Message = err.Error()
		return
	}
	fmt.Printf("request:%+v", request)

	foo, ok := funcList[request.Func]
	if !ok {
		res.ErrorCode = -1
		res.Message = fmt.Sprintf("no func[%s]", request.Func)
		return
	}
	f := reflect.ValueOf(foo)
	paramsArray := request.Params.([]interface{})

	if len(paramsArray) != f.Type().NumIn() {
		panic(fmt.Sprintf("func[%s]  params[%d] not match give[%d]", request.Func, f.Type().NumIn(), len(paramsArray)))
	}

	in := make([]reflect.Value, len(paramsArray))

	for i, v := range paramsArray {
		in[i] = reflect.ValueOf(jsontool.TypeValue(f.Type().In(i), v))
	}
	tres := f.Call(in)
	for k, v := range tres {
		res.Result[fmt.Sprintf("Res_%d", k)] = v.Interface()
	}
}
