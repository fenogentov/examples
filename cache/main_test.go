package main

import (
	"container/ring"
	"strconv"
	"testing"
	"time"
)

var r = []record{}

func init() {
	for i := 0; i < 1024; i++ {
		r = append(r, record{
			Level:      int64(i),
			Time:       time.Now(),
			LoggerName: "test",
			Message:    "Test message # " + strconv.Itoa(i),
		})
	}
}

func BenchmarkMyList(b *testing.B) {
	ram := NewList(512)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1024; j++ {
			ram.Append(r[j])
		}

		for k := 0; k < 10; k++ {
			ram.Get()
		}
	}
}

func BenchmarkRingBuffer(b *testing.B) {
	ram := ring.New(512)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1024; j++ {
			ram.Value = r[j]
			ram = ram.Next()
		}

		for k := 0; k < 10; k++ {
			ram = ram.Next()
		}
	}
}
