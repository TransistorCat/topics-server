package main

import (
	"os"
	"strconv"

	"github.com/TransistorCat/topics-server/cotroller"
	"github.com/TransistorCat/topics-server/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := Init(repository.DefaultOptions); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := cotroller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	r.POST("/community/page/post/topic", func(ctx *gin.Context) {
		title, _ := ctx.GetPostForm("title")
		content, _ := ctx.GetPostForm("content")
		data := cotroller.PublishTopic(title, content)
		ctx.JSON(200, data)
	})
	r.POST("/community/page/post/post", func(ctx *gin.Context) {
		parentIDStr, _ := ctx.GetPostForm("parent_id")
		parentID, _ := strconv.ParseInt(parentIDStr, 10, 64)
		content, _ := ctx.GetPostForm("content")
		data := cotroller.PublishPost(parentID, content)
		ctx.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}

func Init(options repository.Options) error {
	if err := repository.Init(options); err != nil {
		return err
	}
	return nil
}
