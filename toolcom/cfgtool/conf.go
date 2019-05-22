package cfgtool

import (
	"fmt"
	"goc/toolcom/errtool"
	"goc/toolcom/iotool"
	"goc/toolcom/jsontool"
	"io/ioutil"
	"strings"
)

type CfgData struct {
	V map[string]interface{}
}

var confInterface []byte

func New(filePath string) *CfgData {
	//reletive path
	if strings.Index(filePath, "/") != 0 {
		filePath = fmt.Sprintf("%s/%s", iotool.GetCurrentDirectory(), filePath)
	}
	buf, err := ioutil.ReadFile(filePath)
	errtool.Errpanic(err)
	cfg_data := &CfgData{}
	jsontool.Decode(string(buf), &cfg_data.V)
	return cfg_data
}

func (p *CfgData) TakeInt(key string) (int, error) {
	m, ok := p.V[key]
	if !ok {
		return 0, fmt.Errorf("can not find key[%s]", key)
	}
	switch v := m.(type) {
	case int:
		return v, nil
	case float64:
		return int(v), nil
	default:
		return 0, fmt.Errorf("key[%s] is not int", key)
	}
}

func (p *CfgData) TakeString(key string) (string, error) {
	m, ok := p.V[key]
	if !ok {
		return "", fmt.Errorf("can not find key[%s]", key)
	}
	switch v := m.(type) {
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("key[%s] is not string", key)
	}
}

func (p *CfgData) TakeFloat(key string) (float64, error) {
	m, ok := p.V[key]
	if !ok {
		return 0, fmt.Errorf("can not find key[%s]", key)
	}
	switch v := m.(type) {
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("key[%s] is not float", key)
	}
}

func (p *CfgData) TakeBool(key string) (bool, error) {
	m, ok := p.V[key]
	if !ok {
		return false, fmt.Errorf("can not find key[%s]", key)
	}
	switch v := m.(type) {
	case bool:
		return v, nil
	default:
		return false, fmt.Errorf("key[%s] is not bool", key)
	}
}

func (p *CfgData) TakeCfg(key string) (*CfgData, error) {
	m, ok := p.V[key]
	if !ok {
		return nil, fmt.Errorf("can not find key[%s]", key)
	}

	switch v := m.(type) {
	case map[string]interface{}:
		return &CfgData{V: v}, nil
	default:
		return nil, fmt.Errorf("key[%s] is not map", key)
	}
}

func (p *CfgData) TakeJson(key string) (string, error) {
	m, ok := p.V[key]
	if !ok {
		return "", fmt.Errorf("can not find key[%s]", key)
	}
	return jsontool.Encode(m), nil
}