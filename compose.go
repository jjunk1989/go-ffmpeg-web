package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleCompose(c *gin.Context) {
	fmt.Println("start  compose")

	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["files"]
	if len(files) < 2 {
		c.JSON(http.StatusOK, result(RESULT_ERR, "not enough input file", gin.H{}))
		return
	}
	if getFileType(files[0]) != MIME_GIF {
		c.JSON(http.StatusOK, result(RESULT_ERR, "wrong gif file type", gin.H{}))
		return
	}
	if getFileType(files[1]) != MIME_MP3 {
		c.JSON(http.StatusOK, result(RESULT_ERR, "wrong mp3 file type", gin.H{}))
		return
	}
	fmt.Println("file1", files[0].Filename, getFileType(files[0]))
	fmt.Println("file2", files[0].Filename, getFileType(files[1]))
	c.JSON(http.StatusOK, result(RESULT_SUCCESS, "", gin.H{}))
}
