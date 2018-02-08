package main

import (
	"fmt"
	"os/exec"
)

func testCmd() {
	path, err := exec.LookPath("ffmpeg")
	if err != nil {
		fmt.Printf("look path err: %s \r\n", err)
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
	}
	fmt.Printf("cmd out put: \r\n")
	fmt.Printf("%s", out)
}
