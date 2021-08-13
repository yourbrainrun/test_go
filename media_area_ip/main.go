package main

import (
	"bufio"
	"media_area_ip/data"
	"media_area_ip/request"
	"net/http"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	area := data.NewArea()
	go area.GetData("test.txt")
	go ipList(area, &wg)

	wg.Wait()
}

func ipList(area *data.Area, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	requestData := make(map[string]string)
	requestData["stream_url"] = "sdn://wss.wangxiao.eaydu.com/live_ali/x_3_test?key=d0e5159c95e4bd63c93ad467bb14f134&tm=74841c9c"
	requestData["client_id"] = "jack"

	fd, _ := os.Create("output.txt")
	defer func() {
		_ = fd.Close()
	}()

	buf := bufio.NewWriter(fd)

	for v := range area.DataChan {
		client := request.NewBaseService("http://bayer-sdn-api.cloudhub.vip")
		requestData["self_ip"] = v
		str, _ := client.Request(http.MethodPost, "/sdn/subscribe", requestData)
		//fmt.Println(str, err)
		_, _ = buf.WriteString(str + "\n")
	}
}
