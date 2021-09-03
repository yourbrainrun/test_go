package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yourbrainrun/test_go/test_gin_mysql/drivers"
	"github.com/yourbrainrun/test_go/test_gin_mysql/models"
)

func main() {
	gin.DisableConsoleColor()

	//create()
	//test()
	test1()

	//router := gin.New()
	//router.GET("/test", new(controller.Test).Test)
	//_ = router.Run(":888")

}

func test1() {
	db := drivers.GetDatabase()
	var channel models.Channel
	db.First(&channel, "id=?", 4)
	fmt.Println(channel)
}

func test() {
	db := drivers.GetDatabase()
	var channel models.Channel
	channel.ProjectId = 2
	channel.Id = 0
	channel.StreamChannel = "12313123"
	channel.StreamUrl = "url://asdfasdf.com"
	channel.AppId = "2323123"
	channel.Status = 1
	channel.OrganizationId = 2

	db.Where("stream_channel=?", channel.StreamChannel).FirstOrCreate(&channel)
	fmt.Println(channel)
}

func create() {
	var channel models.Channel
	channel.ProjectId = 2
	channel.StreamChannel = "asdfasdfasdfasdf"
	channel.StreamUrl = "asdfasdfasdfadf"
	channel.AppId = "asdfasdfasdff"
	channel.Status = 1
	channel.OrganizationId = 1

	db := drivers.GetDatabase()
	db.Model(true)
	aa := db.Where("stream_channel=?", channel.StreamChannel).FirstOrCreate(&channel)
	fmt.Println(channel)

	var channel1 models.Channel
	//db.Where("id=?", 7).First(&channel1)
	channel1.Id = 10
	db.First(&channel1)
	fmt.Println(channel1)

	fmt.Println(channel)
	fmt.Println(aa)
}
