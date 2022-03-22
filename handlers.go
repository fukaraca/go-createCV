package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"path/filepath"
)

func index(c *gin.Context) {
	symbolic := "/web/img/janedoe.png"
	c.HTML(200, "index.html", gin.H{
		"example":    example,
		"inputCV":    example,
		"photograph": symbolic,
	})
}

func create(c *gin.Context) {
	info := new(Info)
	cvStr := c.PostForm("cv-input-json")
	err := json.Unmarshal([]byte(cvStr), info)
	//parse JSON and unmarshall to Info struct
	if err != nil {
		log.Println("cv input couldn't be unmarshalled", err)
		c.HTML(400, "index.html", gin.H{
			"inputCV": cvStr,
			"created": "Oupss, something's gone wrong.!",
		})
		return
	}
	//resize image and save to temp folder
	imgToBeResized, header, err := c.Request.FormFile("cv-image")
	if err != nil || header.Size == 0 {
		log.Println("image upload failed or no such file:", err)

	} else {
		filePathString := fmt.Sprintf("./web/img/temp/")
		err = resizeAndSave(&imgToBeResized, filePathString, "temp"+filepath.Ext(header.Filename))
		info.photoPath = AddPath() + template.URL("/web/img/temp/temp") + template.URL(filepath.Ext(header.Filename))

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
	//template to buffer
	buff := info.templater()

	//create cv.pdf from buffer
	err = pdfGenerator(buff)
	if err != nil {
		log.Println("pdf couldn't be generated")
	}
	c.HTML(200, "index.html", gin.H{
		"inputCV": cvStr,
		"created": info.Fullname + ".pdf",
		"pdfPath": "/web/dump/CV.pdf",
	})

}
