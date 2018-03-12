package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Id     uint64 `gorm: PRIMARY_KEY`
	Output string `gorm:"type:varchar(255);"`
	// 0 = int, 1 = finish, 2 = running, 3 = err
	Status   uint8
	FileType string `gorm:"type:varchar(255);"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt time.Time
	lock *sync.RWMutex
}

type TaskFunction func(*Task)

func NewTask() *Task {
	var t Task
	t.Id, _ = sf.NextID()
	t.Status = 0
	t.lock = new(sync.RWMutex)
	db.NewRecord(t)
	db.Create(t)
	return &t
}

func (t *Task) Run(f TaskFunction) {
	t.Status = 0
	t.Save()
	go f(t)
}

func (t *Task) Save() {
	t.lock.Lock()
	defer t.lock.Unlock()
	db.Model(t).Save(t)
}

func (t *Task) Delete() {
	db.Delete(t)
}

func handleTask(c *gin.Context) {
	tid := c.Param("tid")
	var task Task
	db.First(&task, tid)
	logan.Info(task)
	if task.Id != 0 {
		c.JSON(http.StatusOK, result(RESULT_SUCCESS, "", gin.H{
			"id":       task.Id,
			"output":   task.Output,
			"status":   task.Status,
			"fileType": task.FileType,
		}))
	} else {
		c.JSON(http.StatusOK, result(RESULT_ERR, "not found task", gin.H{}))
	}
}
