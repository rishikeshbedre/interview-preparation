package main

import (
	"errors"
	"log"
	"sync"
)

type StackOperation interface {
	Push(element string)
	Pop() (string, error)
	IsEmpty() bool
	Show() error
}

type Stack struct {
	mu         sync.RWMutex
	StackSlice []string
}

func (s *Stack) Push(element string) {
	s.mu.Lock()
	s.StackSlice = append(s.StackSlice, element)
	s.mu.Unlock()
}

func (s *Stack) Pop() (string, error) {
	s.mu.Lock()
	if s.IsEmpty() {
		s.mu.Unlock()
		return "", errors.New("Stack empty")
	}
	length := len(s.StackSlice)
	tempElement := s.StackSlice[length-1]
	s.StackSlice = s.StackSlice[:length-1]
	s.mu.Unlock()
	return tempElement, nil
}

func (s *Stack) IsEmpty() bool {
	if len(s.StackSlice) == 0 {
		return true
	}
	return false
}

func (s *Stack) Show() error {
	s.mu.RLock()
	if s.IsEmpty() {
		s.mu.RUnlock()
		return errors.New("Stack empty")
	}
	log.Println(s.StackSlice)
	s.mu.RUnlock()
	return nil
}

func CreateNewStack() *Stack {
	tempStack := Stack{
		mu:         sync.RWMutex{},
		StackSlice: make([]string, 0, 10),
	}
	return &tempStack
}

func main() {
	var stack StackOperation = CreateNewStack()

	err := stack.Show()
	if err != nil {
		log.Println(err)
	}

	stack.Push("hi")
	stack.Push("rishi")

	err = stack.Show()
	if err != nil {
		log.Println(err)
	}

	popElement, err := stack.Pop()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(popElement)
	}

	err = stack.Show()
	if err != nil {
		log.Println(err)
	}

	popElement, err = stack.Pop()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(popElement)
	}

	err = stack.Show()
	if err != nil {
		log.Println(err)
	}

	popElement, err = stack.Pop()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(popElement)
	}
}
