package message

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Id  int64     `json:"id"`
	DT  time.Time `json:"dt"`
	Msg string    `json:"msg"`
}

// PackageMessage is a simple in-memory database
type PackageMessage struct {
	sync.RWMutex
	Messages map[int64]*Message
	nextId   int64
}

func New() *PackageMessage {
	return &PackageMessage{
		Messages: map[int64]*Message{},
	}
}

// Create is creates new message in database.
func (pm *PackageMessage) Create(msg string, dt time.Time) int64 {
	pm.Lock()
	defer pm.Unlock()

	m := &Message{
		Id:  pm.nextId,
		DT:  dt,
		Msg: msg,
	}
	pm.Messages[pm.nextId] = m
	pm.nextId++

	return m.Id
}

// Update is updates message in database by id.
// If no such id exists, an error is returned.
func (pm *PackageMessage) Update(id int64, msg string, dt time.Time) error {
	pm.Lock()
	defer pm.Unlock()

	if _, ok := pm.Messages[id]; !ok {
		return fmt.Errorf("message with id=%d not found", id)
	}

	m := &Message{
		Id:  id,
		DT:  dt,
		Msg: msg,
	}
	pm.Messages[id] = m

	return nil
}

// Get retrieves message from database by id.
// If no such id exists, an error is returned.
func (pm *PackageMessage) Get(id int64) (Message, error) {
	pm.RLock()
	defer pm.RUnlock()

	if m, ok := pm.Messages[id]; ok {
		return *m, nil
	}

	return Message{}, fmt.Errorf("message with id=%d not found", id)
}

// Delete deletes message with the given id.
// If no such id exists, an error is returned.
func (pm *PackageMessage) Delete(id int64) error {
	pm.Lock()
	defer pm.Unlock()

	if _, ok := pm.Messages[id]; !ok {
		return fmt.Errorf("message with id=%d not found", id)
	}

	delete(pm.Messages, id)

	return nil
}

// // DeleteAllTasks deletes all tasks in the store.
// func (ts *TaskStore) DeleteAllTasks() error {
// 	ts.Lock()
// 	defer ts.Unlock()

// 	ts.tasks = make(map[int]Task)
// 	return nil
// }

// // GetAllTasks returns all the tasks in the store, in arbitrary order.
// func (ts *TaskStore) GetAllTasks() []Task {
// 	ts.Lock()
// 	defer ts.Unlock()

// 	allTasks := make([]Task, 0, len(ts.tasks))
// 	for _, task := range ts.tasks {
// 		allTasks = append(allTasks, task)
// 	}
// 	return allTasks
// }

// // GetTasksByTag returns all the tasks that have the given tag, in arbitrary
// // order.
// func (ts *TaskStore) GetTasksByTag(tag string) []Task {
// 	ts.Lock()
// 	defer ts.Unlock()

// 	var tasks []Task

// taskloop:
// 	for _, task := range ts.tasks {
// 		for _, taskTag := range task.Tags {
// 			if taskTag == tag {
// 				tasks = append(tasks, task)
// 				continue taskloop
// 			}
// 		}
// 	}
// 	return tasks
// }

// // GetTasksByDueDate returns all the tasks that have the given due date, in
// // arbitrary order.
// func (ts *TaskStore) GetTasksByDueDate(year int, month time.Month, day int) []Task {
// 	ts.Lock()
// 	defer ts.Unlock()

// 	var tasks []Task

// 	for _, task := range ts.tasks {
// 		y, m, d := task.Due.Date()
// 		if y == year && m == month && d == day {
// 			tasks = append(tasks, task)
// 		}
// 	}

// 	return tasks
// }
