package main

import (
	"fmt"
	"go-project-example/cotroller"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := cotroller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		fmt.Println("Gin 启动失败")
		os.Exit(-1)
	}
}
