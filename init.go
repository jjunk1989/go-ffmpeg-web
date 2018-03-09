package main

import (
	"io/ioutil"
	"os"

	"github.com/sony/sonyflake"
)

// for task id

var sf *sonyflake.Sonyflake
var logan *Logan
var ginLogan *Logan
var tempDir string

func init() {
	// init log
	logan = newLogan("log", "log.log")
	logan.initDir()
	ginLogan = newLogan("log", "access.log")
	ginLogan.initDir()

	// init upload folder
	_, err := os.Stat(UPLOAD_BASE)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(UPLOAD_BASE, 0755)
		} else {
			panic("upload folder not created" + err.Error())
		}
	}
	// int temp dir
	tempDir, err = ioutil.TempDir("", "go_ffmpeg_")
	if err != nil {
		panic("tempDir err" + err.Error())
	}
	// sonwflake setting
	var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}
