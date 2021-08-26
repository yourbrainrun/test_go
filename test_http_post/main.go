package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//err := httpDo()
	err := post2()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func httpDo() error {
	client := &http.Client{}
	urlStr := "http://127.0.0.1:999/sdn/subscribe"
	body := bytes.NewBuffer(nil)

	values := url.Values{}
	values.Set("stream_url", "sdn://wss.wangxiao.eaydu.com/live_ali/x_3_test?key=d0e5159c95e4bd63c93ad467bb14f134&tm=74841c9c")
	values.Set("client_id", "jack")
	body.WriteString(values.Encode())

	req, err := http.NewRequest(http.MethodPost, urlStr, body)
	if err != nil {
		fmt.Println(1)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(2)
		return err
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(3)
		return err
	}
	fmt.Println(string(all))
	return nil
}

//  segment in URL cannot contain colon  没有协议头 会出现
func post2() error {
	urlStr := "http://127.0.0.1:999/sdn/subscribe"
	body := bytes.NewBuffer(nil)

	values := url.Values{}
	values.Set("stream_url", "sdn://wss.wangxiao.eaydu.com/live_ali/x_3_test?key=d0e5159c95e4bd63c93ad467bb14f134&tm=74841c9c")
	values.Set("client_id", "jack")
	body.WriteString(values.Encode())

	resp, err := http.Post(urlStr, "application/x-www-form-urlencoded", body)

	if err != nil {
		fmt.Println(2)
		return err
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(3)
		return err
	}
	fmt.Println(string(all))
	return nil
}
