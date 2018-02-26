package main

import (
	"fmt"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func handleCompose(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	if len(files) < 2 {
		result(RESULT_ERR, "not enough input file", gin.H{})
		return
	}
	fmt.Println(getFileType(files[0]))
	fmt.Println(getFileType(files[1]))
}
