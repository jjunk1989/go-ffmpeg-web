package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleUpload(c *gin.Context) {
	fmt.Println("on upload")
	// single file
	file, _ := c.FormFile("files")
	fmt.Println("file name", file.Filename)
	fmt.Println("flle type", file.Header.Get("Content-Type"))

	// err := c.SaveUploadedFile(file, UPLOAD_BASE+file.Filename)

	// if err != nil {
	//	fmt.Println("save file failed", err)
	// }
	c.JSON(http.StatusOK, result(RESULT_SUCCESS, "", gin.H{"version": "0.1"}))

}
