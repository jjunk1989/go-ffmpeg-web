package main

import (
	"fmt"
	"net/http"
	_ "os"
	"os/exec"
	_ "strconv"
	"sync"
	_ "time"

	"github.com/gin-gonic/gin"
)

func handleTest(c *gin.Context) {
	t := c.Query("type")
	switch t {
	case "version":
		c.JSON(http.StatusOK, result(RESULT_SUCCESS, "", gin.H{"version": "0.1"}))
	case "cmd":
		r := testCmd()
		if r == 0 {
			c.JSON(http.StatusOK, result(RESULT_SUCCESS, "", gin.H{"version": "0.1"}))
		} else {
			c.JSON(http.StatusOK, result(RESULT_ERR, "转换失败", gin.H{"version": "0.1"}))
		}
	default:
		c.JSON(http.StatusOK, result(RESULT_SUCCESS, "unknown cmd", gin.H{"version": "0.1"}))
	}
}

func testCmd() (res int) {
	res = 0
	path, err := exec.LookPath("ffmpeg")
	if err != nil {
		fmt.Printf("look path err: %s \r\n", err)
		res = -1
		return
	}
	fmt.Printf("find path is %s \r\n", path)
	c := exec.Command("ffmpeg",
		"-i", "test.gif",
		"-t", "5",
		"-ss", "00:00:00",
		"-i", "test.mp3",
		"-c:v", "libx264",
		"-c:a", "aac",
		"-b:a", "128k",
		"-vf", "scale=420:-2,format=yuv420p",
		"out.mp4")

	out, err := c.CombinedOutput()
	if err != nil {
		fmt.Printf("run cmd error %s \r\n", err)
		res = -1
	}
	fmt.Printf("cmd out put: \r\n")
	fmt.Printf("%s", out)
	return
}

func testlog() {
	var wg sync.WaitGroup
	lock := new(sync.RWMutex)

	for i := 0; i < 1000; i += 3 {
		go func(i int) {
			defer lock.Unlock()
			defer wg.Done()
			lock.Lock()
			logan.Info("test info", i)
			logan.Warn("test warn", i+1)
			logan.Error("test error", i+2)
		}(i)
		wg.Add(1)
	}
	/*
		err = ioutil.WriteFile("log/test.log", []byte("test log file"), 0755)
		if err != nil {
			fmt.Println("write file err", err)
		}
	*/
	wg.Wait()
}
