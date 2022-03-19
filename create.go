package main

import "github.com/gin-gonic/gin"

func create(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"created": "cv.pdf",
	})
}
