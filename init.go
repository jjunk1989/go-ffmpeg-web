package main

import (
	"os"

	"github.com/sony/sonyflake"
)

// for task id
var sf *sonyflake.Sonyflake
var logan *Logan
var ginLogan *Logan

func init() {
	// init log
	logan = newLogan("log", "log.log")
	logan.initDir()
	ginLogan = newLogan("log", "access.log")
	ginLogan.initDir()

	// init upload folder
	_, err := os.Stat("upload")
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("upload", 0755)
		} else {
			panic("upload folder not created" + err.Error())
		}
	}
	// sonwflake setting
	var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}
