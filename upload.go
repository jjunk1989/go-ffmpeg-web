package main

import (
	_ "fmt"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func handleUpload(c *gin.Context) {
	// fmt.Println("on upload")
	/// c.JSON(http.StatusOK, result(RESULT_SUCCESS, "", gin.H{"version": "0.1"}))
	c.JSON(200, gin.H{"t": 1})
}
