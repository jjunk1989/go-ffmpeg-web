package main

import (
	_ "net/http"

	"github.com/gin-gonic/gin"
)

const (
	RESULT_SUCCESS = iota
	RESULT_ERR
)

func engine() *gin.Engine {

	r := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// r.MaxMultipartMemory = 10 << 20 // 10 MiB

	api := r.Group("api")

	{
		api.GET("test", handleTest)
		api.POST("upload", handleUpload)
	}

	// for web test
	r.Static("/examples", "./examples/web")
	// for upload files
	// TODO unsafe for access!
	// r.Static("/upload", "./upload")
	return r
}

func result(code int, message string, result gin.H) gin.H {
	return gin.H{"code": code, "message": message, "result": result}
}
