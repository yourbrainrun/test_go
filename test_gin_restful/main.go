package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/", Create)
		v1.GET("/:id", Show)
	}

	v2 := router.Group("v2")
	{
		v2.GET("/", func(c *gin.Context) {
			c.JSON(200, map[string]string{
				"v2": "/",
			})
		})

		v2.GET("/test", func(c *gin.Context) {
			c.JSON(200, map[string]string{
				"v2": "/test",
			})
		})
	}

	router.Run(":8088")
}

func Create(c *gin.Context) {
	c.JSON(200, map[string]string{
		"v1": "test",
	})
}
func Show(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, map[string]string{
		"v1": id,
	})
}
