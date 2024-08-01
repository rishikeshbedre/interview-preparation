package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type randomNumber struct {
	num int
	mu  sync.RWMutex
}

func NewRandomNumber() *randomNumber {
	return &randomNumber{
		mu: sync.RWMutex{},
	}
}

func printRandomIntegers(ctx context.Context, randNumber *randomNumber) {
	for {
		select {
		case <-ctx.Done():
			log.Println("exiting print function")
			return
		default:
			log.Println("i am printing random number")
			randNumber.mu.RLock()
			log.Println("interger is ", randNumber.num)
			randNumber.mu.RUnlock()
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	randNumber := NewRandomNumber()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		<-c
		cancel()
	}()

	go printRandomIntegers(ctx, randNumber)

	for {
		select {
		case <-ctx.Done():
			log.Println("context got timeout")
			return
		default:
			log.Println("i am inside generating random number")
			randNumber.mu.Lock()
			randNumber.num = rand.Intn(100)
			randNumber.mu.Unlock()
		}
		time.Sleep(1 * time.Second)
	}
}
