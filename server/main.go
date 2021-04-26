package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Post struct {
	Url  string `json:"url"`
	Name string `json:"name"`
	Ext  string `json:"ext"`
}

func main() {
	storage, ok := os.LookupEnv("AZUKI_STORAGE_PATH")
	if !ok {
		storage = "."
	}

	client := NewClient()
	downloader := NewDownloader(client, storage, 1000, 1000)

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/downloads", func(c *gin.Context) {
		items := downloader.table.ToItems()
		c.JSON(http.StatusOK, items)
	})

	r.POST("/downloads", func(c *gin.Context) {
		var post Post
		c.BindJSON(&post)

		go func() {
			err := downloader.Download(post.Url, post.Name, post.Ext)
			if err != nil {
				fmt.Printf("%+v\n", err)
			}
		}()

		c.JSON(http.StatusCreated, gin.H{})
	})

	r.DELETE("/downloads/:id", func(c *gin.Context) {
		id := c.Param("id")
		downloader.table.Cancel(id)

		c.JSON(http.StatusNoContent, gin.H{})
	})

	r.Run()
}
