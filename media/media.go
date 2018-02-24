package media

import (
	"errors"
	"fmt"
	"os/exec"
)

const CMD_BASE = "ffmpeg"

var installErr = errors.New("Cant find ffmpeg. Please check install and env")

type Media struct {
	Inputs      []string
	Time        string
	StartTime   string
	VideoCode   string
	AudioCode   string
	VideoFormat string
	Bitrate     string
	Output      string
}

func (m Media) CheckSupport() (err error) {
	_, err = exec.LookPath(CMD_BASE)
	if err != nil {
		err = installErr
		return
	}
	return
}

func (m Media) Run() {
	if len(m.Inputs) == 0 || m.Output == "" {
		fmt.Printf("must have input")
		return
	}
	c := exec.Command(CMD_BASE,
		"-i", "test.gif",
		"-i", "test.mp3",
		"-t", m.Time,
		"-ss", m.StartTime,
		"-c:v", m.VideoCode,
		"-c:a", m.AudioCode,
		"-b:a", m.Bitrate,
		"-vf", m.VideoFormat,
		m.Output)
	out, err := c.CombinedOutput()
	if err != nil {
		fmt.Printf("run cmd error %s \r\n", err)
	}
	fmt.Printf("cmd out put: \r\n")
	fmt.Printf("%s", out)
}

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
