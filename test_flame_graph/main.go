package main

import (
	data "github.com/yourbrainrun/test_go/test_flame_graph/pprofDemo"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("Hello World"))
			time.Sleep(1 * time.Second)
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
