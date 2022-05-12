package main

import (
	"sync"
	"time"
)

type record struct {
	Level      int64
	Time       time.Time
	LoggerName string
	Message    string
}

type Item struct {
	Record record
	Prev   *Item
	Next   *Item
}

type Lists struct {
	head     *Item
	tail     *Item
	size     int64
	capacity int64
	item     *Item
	mx       sync.Mutex
}

func NewList(capacity int64) *Lists {
	return &Lists{
		size:     0,
		capacity: capacity,
		mx:       sync.Mutex{},
	}
}

func (l *Lists) Append(r record) {
	if l.size == l.capacity {
		l.removeFront()
	}
	l.addBack(r)
}

func (l *Lists) removeFront() {
	l.mx.Lock()
	defer l.mx.Unlock()

	l.head = l.head.Next
	l.head.Prev = nil
	l.size--
}

func (l *Lists) addBack(r record) {
	l.mx.Lock()
	defer l.mx.Unlock()

	newElem := Item{Record: r, Prev: l.tail, Next: nil}

	if l.size == 0 {
		l.head = &newElem
		l.item = l.head
	} else {
		l.tail.Next = &newElem
	}

	l.tail = &newElem
	l.size++
}

func (l *Lists) Get() record {
	r := l.item.Record
	l.item = l.item.Next
	return r
}

func (l *Lists) MoveToFront() {
	l.item = l.head
}
