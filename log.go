package main

import (
	"fmt"
	_ "io/ioutil"
	"os"
	"time"
)

func log() {
	// os.Mkdir()
	info, err := os.Stat("log")
	if err != nil && os.IsNotExist(err) {
		fmt.Println("Stat file err", err, os.IsNotExist(err))
		os.Mkdir("log", 0755)
	} else {
		fmt.Println("file stat", info.IsDir())
	}

	f, err := os.OpenFile("log/"+"test.log", os.O_APPEND, 0755)
	defer f.Close()
	if err != nil {
		fmt.Println("open log file err", err)
	}
	f.WriteString("test log file" + time.Now().String() + "\r\n")
	/*
		err = ioutil.WriteFile("log/test.log", []byte("test log file"), 0755)
		if err != nil {
			fmt.Println("write file err", err)
		}
	*/
}
