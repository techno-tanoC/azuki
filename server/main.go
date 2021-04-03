package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Url  string `json:"url"`
	Name string `json:"name"`
	Ext  string `json:"ext"`
}

func main() {
	client := NewClient()
	downloader := NewDownloader(client, ".", 1000, 1000)
	logger := log.Default()

	r := gin.Default()

	r.GET("/downloads", func(c *gin.Context) {
		items := downloader.table.ToItems()
		c.JSON(http.StatusOK, items)
	})

	r.POST("/downloads", func(c *gin.Context) {
		var post Post
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		}

		go func() {
			err := downloader.Download(post.Url, post.Name, post.Ext)
			if err != nil {
				logger.Println(err)
			}
		}()

		c.JSON(http.StatusCreated, gin.H{})
	})

	r.DELETE("/downloads/:id", func(c *gin.Context) {
		id := c.Param("id")
		downloader.table.Delete(id)

		c.JSON(http.StatusNoContent, gin.H{})
	})

	r.Run()
}
