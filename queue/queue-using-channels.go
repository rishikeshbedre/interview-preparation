package main

import (
	"errors"
	"log"
	"sync"
)

// Circular Queue using Channels
type QueueOperation interface {
	Enqueue(element string) error
	Dequeue() (string, error)
	IsFull() bool
	IsEmpty() bool
}

type Queue struct {
	QueueChan chan string
	mu        sync.RWMutex
}

func (q *Queue) Enqueue(element string) error {
	q.mu.Lock()
	if q.IsFull() {
		q.mu.Unlock()
		return errors.New("Queue full")
	}
	q.QueueChan <- element
	q.mu.Unlock()
	return nil
}

func (q *Queue) Dequeue() (string, error) {
	q.mu.Lock()
	if q.IsEmpty() {
		q.mu.Unlock()
		return "", errors.New("Queue Empty")
	}
	tempElement, ok := <-q.QueueChan
	q.mu.Unlock()
	if !ok {
		return "", errors.New("Error: channel closed")
	}
	return tempElement, nil
}

func (q *Queue) IsFull() bool {
	if len(q.QueueChan) == cap(q.QueueChan) {
		return true
	}
	return false
}

func (q *Queue) IsEmpty() bool {
	if len(q.QueueChan) == 0 {
		return true
	}
	return false
}

func CreateNewQueue(capacity int) *Queue {
	tempQueue := Queue{
		QueueChan: make(chan string, capacity),
		mu:        sync.RWMutex{},
	}
	return &tempQueue
}

func main() {
	var queue QueueOperation = CreateNewQueue(2)

	dequeueElement, err := queue.Dequeue()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(dequeueElement)
	}

	err = queue.Enqueue("hi")
	if err != nil {
		log.Println(err)
	}

	err = queue.Enqueue("rishi")
	if err != nil {
		log.Println(err)
	}

	err = queue.Enqueue("bye")
	if err != nil {
		log.Println(err)
	}

	dequeueElement, err = queue.Dequeue()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(dequeueElement)
	}

	err = queue.Enqueue("hi once again")
	if err != nil {
		log.Println(err)
	}

	dequeueElement, err = queue.Dequeue()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(dequeueElement)
	}

	dequeueElement, err = queue.Dequeue()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(dequeueElement)
	}

	dequeueElement, err = queue.Dequeue()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(dequeueElement)
	}
}
