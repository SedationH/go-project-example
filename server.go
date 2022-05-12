package main

import (
	"encoding/json"
	"go-project-example/cotroller"
	"go-project-example/repository"
	"go-project-example/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := Init("./data/"); err != nil {
		os.Exit(-1)
	}

	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		pageData := cotroller.QueryPageInfo(topicId)
		c.JSON(http.StatusOK, pageData)
	})
	r.POST("/community/page/post", func(c *gin.Context) {
		b, err := c.GetRawData()
		if err != nil {
			log.Println("GetRawData() failed")
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		var pageInfo service.PageInfo
		err = json.Unmarshal(b, &pageInfo)
		if err != nil {
			log.Println("Unmarshal() failed", string(b))
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		pageData := cotroller.AddNewPage(&pageInfo)
		c.JSON(http.StatusOK, pageData)
	})

	err := r.Run()
	if err != nil {
		log.Println("Gin 启动失败")
		os.Exit(-1)
	}
}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		return err
	}
	return nil
}
