package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handleCompose(c *gin.Context) {
	logan.Info("start  compose")

	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["files"]
	if len(files) < 2 {
		c.JSON(http.StatusOK, result(RESULT_ERR, "not enough input file", gin.H{}))
		return
	}
	if getFileType(files[0]) != MIME_GIF {
		logan.Info("file1", files[0].Filename, getFileType(files[0]))
		c.JSON(http.StatusOK, result(RESULT_ERR, "wrong gif file type", gin.H{}))
		return
	}
	if getFileType(files[1]) != MIME_MP3 {
		logan.Info("file2", files[0].Filename, getFileType(files[1]))
		c.JSON(http.StatusOK, result(RESULT_ERR, "wrong mp3 file type", gin.H{}))
		return
	}

	t := NewTask()
	t.addLocalFileName(UPLOAD_BASE + strconv.FormatUint(t.id, 10) + ".gif")
	t.addLocalFileName(UPLOAD_BASE + strconv.FormatUint(t.id, 10) + ".mp3")

	if err := c.SaveUploadedFile(files[0], t.localFiles[0]); err != nil {
		c.JSON(http.StatusOK, result(RESULT_ERR, "save gif err"+err.Error(), gin.H{}))
		return
	}
	if err := c.SaveUploadedFile(files[1], t.localFiles[1]); err != nil {
		c.JSON(http.StatusOK, result(RESULT_ERR, "save mp3 err"+err.Error(), gin.H{}))
		return
	}
	time := c.DefaultPostForm("time", "5")
	startTime, _ := strconv.Atoi(c.DefaultPostForm("startTime", "0"))
	com := Compose{
		InputGIF:    t.localFiles[0],
		InputMp3:    t.localFiles[1],
		Time:        time,
		StartTime:   formatTimeString(startTime),
		VideoCode:   "libx264",
		AudioCode:   "aac",
		VideoFormat: "scale=420:-2,format=yuv420p",
		Bitrate:     "128k",
		Output:      UPLOAD_BASE + strconv.FormatUint(t.id, 10) + ".mp4",
	}
	if err := com.Run(); err != nil {
		c.JSON(http.StatusOK, result(RESULT_ERR, "compose err"+err.Error(), gin.H{}))
		return
	}
	c.JSON(http.StatusOK, result(RESULT_SUCCESS, "", gin.H{"video": strconv.FormatUint(t.id, 10) + ".mp4"}))
}
