package main

import (
	"fmt"
	"sync"
	"time"
)

// Write a program to simulate a race condition occurring when one goroutine updates a data variable with odd numbers,
// while another updates the same data variable with even numbers.
// After each update , attempt to display the data contained in the data variable to screen.
// [Goroutines][Concurrency][Race Conditions]
func main() {

	var data int
	var mutex sync.Mutex
	var wg sync.WaitGroup // Create a WaitGroup
	dataChannel := make(chan bool, 1)

	wg.Add(2) // We have two goroutines to wait for
	go func() {
		defer wg.Done() // Signal that this goroutine is done
		for i := 0; i <= 10; i += 2 {
			time.Sleep(time.Millisecond * 10)
			mutex.Lock()
			data = i
			fmt.Println(data)
			mutex.Unlock()
		}
		dataChannel <- true
	}()
	go func() {
		defer wg.Done() // Signal that this goroutine is done
		<-dataChannel   //
		for i := 1; i <= 10; i += 2 {
			time.Sleep(time.Millisecond * 10)
			mutex.Lock()
			data = i
			fmt.Println(data)
			mutex.Unlock()
		}
	}()
	//need to have a sleep here as the executions won't have finished by the time main is done, therefore the print statements won't print aas they might be in flight.
	// this would be alternative to waitgroup
	//time.Sleep(2 * time.Second)
	wg.Wait() // Wait for all goroutines to finish

	// Create a buffered channel with a capacity of 1
	ch := make(chan bool, 1)

	// Send an initial value into the channel
	ch <- true

	// Receive the value from the channel
	value := <-ch
	fmt.Println("Received:", value)
}
