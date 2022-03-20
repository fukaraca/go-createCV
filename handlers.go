package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"image"
	"log"
)

func index(c *gin.Context) {

	c.HTML(200, "index.html", gin.H{
		"example": example,
		"inputCV": example,
	})
}

func create(c *gin.Context) {
	info := new(Info)
	img := new(image.NRGBA)
	cvStr := c.PostForm("cv-input-json")
	err := json.Unmarshal([]byte(cvStr), info)
	if err != nil {
		log.Println("cv input couldn't be unmarshalled", err)
		c.HTML(400, "index.html", gin.H{
			"inputCV": cvStr,
			"created": "Oupss, something's gone wrong.!",
		})
		return
	}
	imgToBeResized, header, err := c.Request.FormFile("cv-image")
	if err != nil || header.Size == 0 {
		log.Println("image upload failed or no such file:", err)
	} else {
		filename := info.Fullname
		filePathString := fmt.Sprintf("./web/img/temp/")
		err, img = resizer(&imgToBeResized, filePathString, filename)
		if err != nil {
			log.Println("image couldn't be resized", err)
			c.HTML(400, "index.html", gin.H{
				"inputCV": cvStr,
				"created": "Oupss, something's gone wrong.!",
			})
			return
		}
		defer imgToBeResized.Close()
	}
	//todo
	*img = image.NRGBA{}

	c.HTML(200, "index.html", gin.H{
		"message": "pdf was created succesfully",
		"inputCV": cvStr,
	})

}
