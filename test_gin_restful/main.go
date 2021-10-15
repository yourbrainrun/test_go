package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	v1 := router.Group("/api/vi/user")
	{
		v1.POST("/", Create)
		v1.POST("/:id", Show)
	}
}

func Create(c *gin.Context) {
	c.JSON(200, map[string]string{
		"test": "test",
	})
}
func Show(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, map[string]string{
		"test": id,
	})
}
