package main

import (
	_ "github.com/gin-gonic/gin"
)

type Task struct {
	id         uint64
	localFiles []string
	output     string
}

func NewTask() *Task {
	var t Task
	t.id, _ = sf.NextID()
	t.localFiles = make([]string, 0)
	return &t
}

func (t *Task) addLocalFileName(name string) {
	t.localFiles = append(t.localFiles, name)
}
