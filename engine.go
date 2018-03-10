package main

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
)

func engine() *gin.Engine {
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(ginLogan.File, os.Stdout)

	r := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 10 << 20 // 10 MiB

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	api := r.Group("api")

	{
		api.GET("test", handleTest)
		api.GET("task/:tid", handleTask)
		api.POST("upload", handleUpload)
		api.POST("compose", handleCompose)
		api.POST("converter", handleConverter)
	}

	// for web test
	r.Static("/examples", "./examples/web")
	// for upload files
	// TODO unsafe for access!
	r.Static("/upload", "./upload")
	return r
}

func result(code int, message string, result gin.H) gin.H {
	return gin.H{"code": code, "message": message, "result": result}
}

func getFileType(file *multipart.FileHeader) string {
	return file.Header.Get(FILE_CONTENT_TYPE)
}
