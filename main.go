package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	ready := false

	go func() {
		fmt.Println("Goroutine: Waiting for the condition...")

		mutex.Lock()
		for !ready {
			cond.Wait()
		}
		fmt.Println("Goroutine: Condition met, proceeding...")
		mutex.Unlock()
	}()

	time.Sleep(2 * time.Second)

	mutex.Lock()
	ready = true
	cond.Signal()
	mutex.Unlock()
	fmt.Println("Push signal !")

	time.Sleep(1 * time.Second)
	fmt.Println("Main: Work is done.")
}
