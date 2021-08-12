package main

import (
	"carina-exporter/collectors"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	str := `{
    "time": "2021-08-11 12:01:43",
    "module": "sdn_server",
    "biz_category": "sdn",
    "msg_name": "link_close",
    "error_code": 0,
    "error_tips": "remove",
    "room_id": "haoweilai_sdn://wss.wangxiao.eaydu.com/live/test",
    "stream_id": "8888",
    "sdn_link_info": {
        "created_time": "2021-08-11 11:02:04",
        "my_lan_ip": "172.16.0.9",
        "my_wan_ip": "123.56.227.82",
        "my_port": 5816,
        "peer_ip": "59.110.66.6",
        "peer_port": 1802,
        "protocol": "udp",
        "direction": "recv"
    }
}`
	collectors.PutSlsLog(str)
}
