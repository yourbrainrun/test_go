package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Test struct {
}

func (*Test) Test(c *gin.Context) {
	name, _ := c.Get("name")
	fmt.Println(name)
}
