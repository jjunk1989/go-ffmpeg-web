package main

import (
	"fmt"
	_ "io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"
)

func testlog() {
	// os.Mkdir()
	info, err := os.Stat("log")
	if err != nil && os.IsNotExist(err) {
		fmt.Println("Stat file err", err, os.IsNotExist(err))
		os.Mkdir("log", 0755)
	} else {
		fmt.Println("file stat", info.IsDir())
	}

	f, err := os.OpenFile("log/"+"test.log", os.O_APPEND|os.O_CREATE, 0755)
	defer f.Close()
	if err != nil {
		fmt.Println("open log file err", err)
	}
	f.WriteString("test log file" + time.Now().String() + "\r\n")

	var wg sync.WaitGroup
	lock := new(sync.RWMutex)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer lock.Unlock()
			defer wg.Done()
			lock.Lock()
			f.WriteString("test log file :" + strconv.Itoa(i) + ";time:" + time.Now().String() + "\r\n")
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
