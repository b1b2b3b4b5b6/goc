package jsontool

import (
	"encoding/json"
	"fmt"
	"goc/toolcom/dumptool"
	"goc/toolcom/errtool"
	"reflect"
	"strings"
)

func DecodeString(jsonStr string, key string) (string, error) {
	var m interface{}
	Decode(jsonStr, &m)

	t := m.(map[string]interface{})
	v, ok := t[key]
	if !ok {
		panic(fmt.Sprintf("can not find %s", key))
	}
	return v.(string), nil
}

func DecodeInt(jsonStr string, key string) (int, error) {
	var m interface{}
	Decode(jsonStr, &m)

	t := m.(map[string]interface{})
	v, ok := t[key]
	if !ok {
		panic(fmt.Sprintf("can not find %s", key))
	}
	return int(v.(float64)), nil
}

func Encode(arg interface{}) string {
	json_byte, err := json.Marshal(arg)
	errtool.Errpanic(err)
	return string(json_byte)
}

func EncodeIndent(arg interface{}) string {
	json_byte, err := json.MarshalIndent(arg, "", "\t")
	errtool.Errpanic(err)
	return string(json_byte)
}

func Decode(jsonStr string, arg interface{}) error {
	strings.TrimRight(jsonStr, "\x00")
	err := json.Unmarshal([]byte(jsonStr), arg)
	errtool.Errpanic(err)
	return err
}

func TypeValue(typ reflect.Type, value interface{}) interface{} {
	switch value.(type) {
	case float64:
		v := value.(float64)
		t := reflect.New(typ)
		switch t.Interface().(type) {
		case *int:
			return int(v)
		case *int16:
			return int16(v)
		case *int32:
			return int32(v)
		case *int64:
			return int64(v)

		case *uint:
			return uint(v)
		case *uint16:
			return uint16(v)
		case *uint32:
			return uint32(v)
		case *uint64:
			return uint64(v)

		case *float32:
			return float32(v)
		case *float64:
			return float64(v)
		default:
			panic(fmt.Sprintf("value[%+v] can not convert to typ[%+v] ", dumptool.Sdump(value), typ))
			return nil
		}
	case string:
		return value

	case bool:
		return value

	default:
		panic(fmt.Sprintf("json not no support [%s]", dumptool.Sdump(value)))
		return nil
	}
}
