package main

import (
	"errors"
	"os/exec"
)

const CMD_BASE = "ffmpeg"

var installErr = errors.New("Cant find ffmpeg. Please check install and env")
var inputErr = errors.New("must have input")
var cmdErr = errors.New("run cmd error")

type Compose struct {
	InputGIF    string
	InputMp3    string
	Time        string
	StartTime   string
	VideoCode   string
	AudioCode   string
	VideoFormat string
	Bitrate     string
	Output      string
}

func (m Compose) CheckSupport() (err error) {
	_, err = exec.LookPath(CMD_BASE)
	if err != nil {
		err = installErr
		return
	}
	return
}

func (m Compose) Run() error {
	if m.InputGIF == "" || m.InputMp3 == "" || m.Output == "" {
		logan.Error("must have input")
		return inputErr
	}

	c := exec.Command(CMD_BASE,
		"-i", m.InputGIF,
		"-t", m.Time,
		"-ss", m.StartTime,
		"-i", m.InputMp3,
		"-c:v", m.VideoCode,
		"-c:a", m.AudioCode,
		// [aac @ 0x22b1a80] The encoder 'aac' is experimental but experimental
		// codecs are not enabled,
		// add '-strict -2' if you want to use it.
		"-strict", "-2",
		"-b:a", m.Bitrate,
		"-vf", m.VideoFormat,
		m.Output)
	out, err := c.CombinedOutput()
	if err != nil {
		logan.Error("run cmd error", err)
		return cmdErr
	}
	logan.Info("cmd out put: \r\n")
	logan.Printf("%s", out)
	return nil
}
