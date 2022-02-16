package main

import (
	"fmt"
	"google.golang.org/grpc/status"
)

func main() {
	err := ret()

	code := status.Code(err)
	msg := status.Convert(err).Message()
	if code == 100001 {
		fmt.Println(code, msg)
	} else {
		fmt.Println("no")
	}

}

func ret() error {
	return status.Error(100001, "{\"code\":123}")
}
