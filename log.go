package main

import (
	"fmt"
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
	logger   *log.Logger
}

func newLogan(path string, name string) *Logan {
	return &Logan{
		Path:     path,
		FileName: name,
		lock:     new(sync.RWMutex),
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
	if err != nil {
		return
	}
	l.logger = log.New(l.File, "", log.Ldate|log.Ltime|log.Lshortfile)
	return
}

// short for close file
func (l *Logan) close() (err error) {
	l.Info("close log file")
	err = l.File.Close()
	return
}

func (l *Logan) Info(v ...interface{}) {
	l.logger.SetPrefix("[INFO]: ")
	l.logger.Output(2, fmt.Sprintln(v...))
}
func (l *Logan) Warn(v ...interface{}) {
	l.logger.SetPrefix("[WARN]: ")
	l.logger.Output(2, fmt.Sprintln(v...))
}

func (l *Logan) Error(v ...interface{}) {
	l.logger.SetPrefix("[ERROR]:")
	l.logger.Output(2, fmt.Sprintln(v...))
}

func (l *Logan) Panic(v ...interface{}) {
	l.logger.SetPrefix("[PANIC]:")
	l.logger.Panicln(v...)
}

func (l *Logan) Printf(format string, v ...interface{}) {
	l.logger.SetPrefix("")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}
