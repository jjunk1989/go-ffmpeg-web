package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

	"sync"
)

type Logan struct {
	Path     string
	FileName string
	File     *os.File
	lock     *sync.RWMutex
	buf      *bytes.Buffer
	logger   *log.Logger
}

func newLogan() *Logan {
	var buf bytes.Buffer

	return &Logan{
		Path:     "log",
		FileName: "log.log",
		lock:     new(sync.RWMutex),
		buf:      &buf,
		logger:   log.New(&buf, "", log.Ldate|log.Ltime),
	}
}

func (l Logan) initDir() (err error) {
	_, err = os.Stat(l.Path)

	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(l.Path, 0755)
		} else {
			return
		}
	}
	return
}

// short for Openfile
func (l *Logan) open() (err error) {
	l.File, err = os.OpenFile(filepath.Join(l.Path, l.FileName), os.O_APPEND|os.O_CREATE, 0755)
	return
}

// short for close file
func (l *Logan) close() (err error) {
	err = l.File.Close()
	return
}

func (l *Logan) Info(v ...interface{}) {
	l.logger.SetPrefix("[INFO]: ")
	l.logger.Println(v...)
	l.write()
}
func (l *Logan) Warn(v ...interface{}) {
	l.logger.SetPrefix("[WARN]: ")
	l.logger.Println(v...)
	l.write()
}

func (l *Logan) Error(v ...interface{}) {
	l.logger.SetPrefix("[ERROR]:")
	l.logger.Println(v...)
	l.write()
}

func (l *Logan) Panic(v ...interface{}) {
	l.logger.SetPrefix("[PANIC]:")
	l.logger.Panicln(v...)
}

func (l *Logan) Printf(format string, v ...interface{}) {
	l.logger.SetPrefix("")
	l.logger.Printf(format, v...)
	l.write()
}

func (l *Logan) write() {
	defer l.lock.Unlock()
	l.lock.Lock()
	l.File.Write(l.buf.Bytes())
	l.buf.Reset()
}
