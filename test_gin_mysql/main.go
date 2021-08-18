package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yourbrainrun/test_go/test_gin_mysql/drivers"
	"github.com/yourbrainrun/test_go/test_gin_mysql/models"
)

func main() {
	gin.DisableConsoleColor()

	create()

	//router := gin.New()
	//router.GET("/test", new(controller.Test).Test)
	//_ = router.Run(":888")

}

func create() {
	var channel models.Channel
	channel.ProjectId = 12
	channel.StreamChannel = "asdfasdfasdfasdf"
	channel.StreamUrl = "asdfasdfasdfadf"
	channel.AppId = "asdfasdfasdff"
	channel.Status = 1
	channel.OrganizationId = 1

	aa := drivers.GetDatabase().FirstOrCreate(&channel)

	fmt.Println(channel)
	fmt.Println(aa)
}
