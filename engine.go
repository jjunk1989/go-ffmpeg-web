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

	api := r.Group("api")

	{
		api.GET("test", handleTest)
		api.POST("upload", handleUpload)
	}
	// for web test
	r.Static("/examples", "./examples/web")

	return r
}

func result(code int, message string, result gin.H) gin.H {
	return gin.H{"code": code, "message": message, "result": result}
}