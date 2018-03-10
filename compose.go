package main

import (
	"log"
	"net/http"
	"path/filepath"
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
	log.Println("taskid NewTask", t)
	t.FileType = "mp4"
	t.Output = strconv.FormatUint(t.Id, 10) + ".mp4"
	t.Save()
	log.Println("taskid NewTask Save", t)

	gFile := filepath.Join(tempDir, strconv.FormatUint(t.Id, 10)+".gif")
	mFile := filepath.Join(tempDir, strconv.FormatUint(t.Id, 10)+".mp3")

	if err := c.SaveUploadedFile(files[0], gFile); err != nil {
		c.JSON(http.StatusOK, result(RESULT_ERR, "save gif err"+err.Error(), gin.H{}))
		return
	}
	if err := c.SaveUploadedFile(files[1], mFile); err != nil {
		c.JSON(http.StatusOK, result(RESULT_ERR, "save mp3 err"+err.Error(), gin.H{}))
		return
	}
	time := c.DefaultPostForm("time", "5")
	startTime, _ := strconv.Atoi(c.DefaultPostForm("startTime", "0"))

	log.Println("taskid", t)

	t.Run(func(task *Task) {
		log.Println("taskid NewTask Save", task)
		com := Compose{
			InputGIF:    gFile,
			InputMp3:    mFile,
			Time:        time,
			StartTime:   formatTimeString(startTime),
			VideoCode:   "libx264",
			AudioCode:   "aac",
			VideoFormat: "scale=420:-2,format=yuv420p",
			Bitrate:     "128k",
			Output:      filepath.Join(UPLOAD_BASE, t.Output),
		}

		if err := com.Run(); err != nil {
			task.Status = 3
			task.Save()
			logan.Info("task err", task.Id)
			return
		}
		logan.Info("task finish", task.Id)
		task.Status = 1
		task.Save()
	})

	log.Println("taskid return", t, t.Id)
	c.JSON(http.StatusOK, result(RESULT_SUCCESS, "", gin.H{"task": strconv.FormatUint(t.Id, 10)}))
}
