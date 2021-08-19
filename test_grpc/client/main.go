package main

import (
	"context"
	"fmt"
	my "github.com/yourbrainrun/test_go/test_grpc/proto"
	"google.golang.org/grpc"
)

const netIp = "127.0.0.1:11111"

func main() {
	conn, err := grpc.Dial(netIp, grpc.WithInsecure())
	if err != nil {
		fmt.Println("dial errors")
		return
	}

	defer conn.Close()

	c := my.NewHelloServerClient(conn)
	ret, err := c.SayHello(context.Background(), &my.HelloRequest{
		Name: "hua hua",
	})
	if err != nil {
		fmt.Println("say hello  errors")
		return
	} else {
		fmt.Println("return say hello ", ret.Message)
	}
}
