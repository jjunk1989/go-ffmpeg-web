package main

import (
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sony/sonyflake"
)

// for task id

var (
	sf       *sonyflake.Sonyflake
	logan    *Logan
	ginLogan *Logan
	tempDir  string
	db       *gorm.DB
)

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
	if tempDir, err = ioutil.TempDir("", "go_ffmpeg_"); err != nil {
		panic("tempDir err" + err.Error())
	}
	// sonwflake setting
	var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
	// open db
	// connect to local db. unsafe
	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_ffmpeg?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("open mysql failed: " + err.Error())
	}
}
