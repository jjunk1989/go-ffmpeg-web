package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleUpload(c *gin.Context) {
	fmt.Println("on upload")
	// single file
	file, _ := c.FormFile("gif")
	fmt.Println("file name", file.Filename)
	fmt.Println("flle type", file.Header.Get("Content-Type"))

	err := c.SaveUploadedFile(file, "./upload/"+file.Filename)

	if err != nil {
		fmt.Println("save file failed", err)
	}
	c.JSON(http.StatusOK, result(RESULT_SUCCESS, "", gin.H{"version": "0.1"}))

}
