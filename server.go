package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		c.JSON(200, gin.H{
			"message": topicId,
		})
	})
	err := r.Run()
	if err != nil {
		fmt.Println("Gin 启动失败")
		os.Exit(-1)
	}
}
