package main

import (
	"context"
	"fmt"
	my "github.com/yourbrainrun/test_go/test_grpc/proto"
	"google.golang.org/grpc"
	"net"
)

const netIp = "127.0.0.1:11111"

type server struct {
}

func (this *server) SayHello(ctx context.Context, in *my.HelloRequest) (*my.HelloReplay, error) {
	return &my.HelloReplay{Message: "say ok" + in.Name}, nil
}

func (this *server) GetHelloMsg(ctx context.Context, in *my.HelloRequest) (*my.HelloMessage, error) {
	return &my.HelloMessage{Msg: "msg say ok" + in.Name}, nil
}

func main() {
	lTcp, err := net.Listen("tcp", netIp)
	if err != nil {
		fmt.Println("network error")
		return
	}
	srv := grpc.NewServer()
	my.RegisterHelloServerServer(srv, &server{})

	err = srv.Serve(lTcp)
	if err != nil {
		fmt.Println("network server error")
		return
	}
}
