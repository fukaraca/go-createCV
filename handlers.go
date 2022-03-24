package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
)

//index represents handler for landing page
func index(c *gin.Context) {
	symbolic := "/web/img/janedoe.png"
	c.HTML(200, "index.html", gin.H{
		"example":    example,
		"inputCV":    example,
		"photograph": symbolic,
	})
}

//create is handler for generating PDF with provided JSON formatted information
func create(c *gin.Context) {
	info := new(Info)
	cvStr := c.PostForm("cv-input-json")
	err := json.Unmarshal([]byte(cvStr), info)
	//parse JSON and unmarshall to Info struct
	if err != nil {
		log.Println("cv input couldn't be unmarshalled", err)
		c.HTML(400, "index.html", gin.H{
			"inputCV": cvStr,
			"message": "Oupss, something's gone wrong.!",
		})
		return
	}
	//resize image and save to temp folder
	imgToBeResized, header, err := c.Request.FormFile("cv-image")
	if err != nil || header.Size == 0 {
		log.Println("image upload failed or no such file:", err)
		info.photoPath = "data:image/gif;base64,R0lGODlhAQABAPcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACH5BAEAAP8ALAAAAAABAAEAAAgEAP8FBAA7"
	} else {
		filePathString := fmt.Sprintf("./web/img/temp/")
		err = resizeAndSave(&imgToBeResized, filePathString, "temp"+filepath.Ext(header.Filename))
		info.photoPath = AddPath() + template.URL("/web/img/temp/temp") + template.URL(filepath.Ext(header.Filename))
		defer func(ext string) {
			err = os.Remove("./web/img/temp/temp" + filepath.Ext(header.Filename))
			if err != nil {
				log.Println("image couldn't be deleted:", err)
			}
		}(filepath.Ext(header.Filename))

		if err != nil {
			log.Println("image couldn't be resized", err)
			c.HTML(400, "index.html", gin.H{
				"inputCV": cvStr,
				"created": "Oupss, something's gone wrong.! Please check JSON format.",
			})
			return
		}
		defer imgToBeResized.Close()
	}
	fileName := info.Fullname + "'s CV.pdf"
	//template to buffer
	buff := info.templater()

	//create cv.pdf from buffer
	err, reader := pdfGenerator(buff)
	if err != nil {
		log.Println("pdf couldn't be generated")
		c.HTML(400, "index.html", gin.H{
			"inputCV": cvStr,
			"created": "Oupss, something's gone wrong.! Pdf couldn't be generated due to some internal errors.:" + err.Error(),
		})
		return
	}
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	_, err = io.Copy(c.Writer, reader)
	if err != nil {
		log.Println("file couldn't be served to client:", err)
	}

}
