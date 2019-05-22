package nettool

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func Http_Post(url string, body string) ([]byte, error) {

	//创建请求
	postReq, err := http.NewRequest("POST",
		url,                     //post链接
		strings.NewReader(body)) //post内容

	if err != nil {
		return nil, err
	}

	//增加header
	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	//执行请求
	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		return nil, err
	}

	//读取响应
	resBody, err := ioutil.ReadAll(resp.Body) //此处可增加输入过滤
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resBody, nil
}
