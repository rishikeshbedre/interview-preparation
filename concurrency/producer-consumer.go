/*
Program to print sum of square of the numbers:
input:
arr = [1,2,3,4,5,6,7,8,9]
n = 4

n is the number of goroutines to be created for generating squares in random order of arr.
create a consumer to receive this squares and sum it up.
main function should recieve this final sum and print it.
use channels where ever necessary.
*/

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func producer(ctx context.Context, prodChan, consumerChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			log.Println("Producer is exiting due to context done")
			return
		case element, isOpen := <-prodChan:
			if !isOpen {
				log.Println("Producer channel closed, exiting producer")
				return
			}
			time.Sleep(5 * time.Second)
			if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
				log.Println("Producer is exiting due to context done")
				return
			}
			consumerChan <- element * element
		}
	}
}

func consumer(ctx context.Context, consumerChan chan int, wg *sync.WaitGroup, sumChan chan int) {
	defer wg.Done()
	sum := 0

	for {
		select {
		case <-ctx.Done():
			log.Println("Consumer is exiting due to context done")
			return
		case square, isOpen := <-consumerChan:
			if !isOpen {
				log.Println("Consumer channel closed, exiting consumer")
				sumChan <- sum
				return
			}
			sum = sum + square
		}
	}
}

func main() {
	osChan := make(chan os.Signal)
	signal.Notify(osChan, os.Interrupt, syscall.SIGTERM)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	nGoRoutines := 4

	prodChan := make(chan int, nGoRoutines)
	consumerChan := make(chan int, nGoRoutines)
	sumChan := make(chan int)

	ctx, cancelFunc := context.WithCancel(context.Background())
	wgProducer := &sync.WaitGroup{}
	wgConsumer := &sync.WaitGroup{}

	go func() {
		<-osChan
		cancelFunc()
		close(sumChan)
	}()

	for i := 0; i < nGoRoutines; i++ {
		go producer(ctx, prodChan, consumerChan, wgProducer)
	}
	wgProducer.Add(nGoRoutines)

	go consumer(ctx, consumerChan, wgConsumer, sumChan)
	wgConsumer.Add(1)

	for _, element := range arr {
		select {
		case <-ctx.Done():
			log.Println("context canceled exiting to send elements to producer")
			break
		case prodChan <- element:
		}
	}
	close(prodChan)

	wgProducer.Wait()
	close(consumerChan)

	log.Println("Sum of squares of the numbers ", <-sumChan)
	wgConsumer.Wait()
	log.Println("Closed all producer and consumer goroutines exiting program successfully")
}
