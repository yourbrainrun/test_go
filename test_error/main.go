package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/status"
)

func main() {
	err := ret()

	type Data struct {
		Code int `json:"code"`
	}
	var badData Data
	code := status.Code(err)
	msg := status.Convert(err).Message()
	if code == 100001 {
		buf := bytes.NewBuffer(nil)
		buf.WriteString(status.Convert(err).Message())
		_ = json.Unmarshal(buf.Bytes(), &badData)
		fmt.Println(badData, "bd", string(buf.Bytes()))
		fmt.Println(msg, "msg")
	} else {
		fmt.Println("no")
	}

}

func ret() error {
	return status.Error(100001, "{\"code\":123}")
}
