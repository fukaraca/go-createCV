package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

/*var Server_Host = os.Getenv("hostAdr")
var Server_Port = os.Getenv("PORT")*/

var Server_Host = "localhost"
var Server_Port = "8080"
var R *gin.Engine
var SecureTag = os.Getenv("IS_SECURE")
var secure bool
