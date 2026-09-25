package main

import (
	"fmt"
	"time"
)

func main() {

	ch := time.Tick(2 * time.Second)

	go func() {
		for range ch {
			fmt.Println("Hello, World!")
		}
	}()

	time.Sleep(6 * time.Second)
}
