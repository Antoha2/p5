package test

import (
	"testing"
	"time"
)

func MyFunction(ch chan<- string) {
	ch <- "Результат"
}

func TestMyConcurrentFunction(t *testing.T) {
	ch := make(chan string)

	go MyFunction(ch)

	select {
	case result := <-ch:
		if result != "Результат" {
			t.Errorf("Unexpected result: %s", result)
		}
	case <-time.After(time.Second):
		t.Error("Timeout waiting for result")
	}
}
