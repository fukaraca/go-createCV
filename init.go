package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

//InitServer function initiates general stuff
func init() {

	//logger middleware teed to log.file
	logfile, err := os.OpenFile("./logs/log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Could not create/open log file")
	}
	errlogfile, err := os.OpenFile("./logs/errlog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Could not create/open err log file")
	}
	gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(errlogfile, os.Stdout)
	//starts with builtin Logger() and Recovery() middlewares
	R = gin.Default()

	//Serving HTML files
	R.LoadHTMLGlob("index.html")

	//Hint: Designated custom Filesystem. When HTML file asks, static.serve shows "root" to look at. indexes:  enables or
	//disables the listing of contents in folder "root"
	//prefix simply strips prefix from each static items path otherwise we will likely see 404 error
	R.Use(static.Serve("/web", static.LocalFile("./web", false)))
	//Heroku thing
	if SecureTag == "true" {
		secure = true
	} else {
		secure = false
	}
}
