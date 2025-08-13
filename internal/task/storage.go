package task

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Data struct {
	tasks map[string]*Task
	mutex *sync.RWMutex
}

func NewData() *Data {
	return &Data{make(map[string]*Task, 20), new(sync.RWMutex)}
}

func (d *Data) AddTask(id string, status string, title string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	CreateTime := time.Now().Format("2006-01-02 15:04:05")
	newTask := &Task{Id: id, Status: status, TimeCreated: CreateTime, TitleTask: title}
	if _, ok := d.tasks[id]; ok {
		return fmt.Errorf("already exists")
	}
	d.tasks[id] = newTask
	return nil
}

func (d *Data) TakeTasks() []Task {
	result := make([]Task, 0, len(d.tasks))
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	for _, v := range d.tasks {
		result = append(result, *v)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].TimeCreated < result[j].TimeCreated
	})
	return result
}

func (d *Data) TakeTaskByID(id string) (*Task, error) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	if v, ok := d.tasks[id]; !ok {
		return nil, fmt.Errorf("not exists")
	} else {
		return v, nil
	}
}

type Storage interface {
	TakeTaskByID(id string) (*Task, error)
	TakeTasks() []Task
	AddTask(id string, status string, title string) error
}

type DataStorage struct {
	Storage
}
