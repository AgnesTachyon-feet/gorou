package main

import (
	"fmt"
)

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go func() {
		channel1 <- 10
		close(channel1)
	}()

	go func() {
		channel2 <- 20
		close(channel2)
	}()

	ClosedChannel1, ClosedChannel2 := false, false
	for {
		if ClosedChannel1 && ClosedChannel2 {
			break
		}
		select {
		case v, ok := <-channel1:
			if !ok {
				ClosedChannel1 = true
				continue
			}
			fmt.Println("Channel1", v)
		case v, ok := <-channel2:
			if !ok {
				ClosedChannel2 = true
				continue
			}
			fmt.Println("Channel2", v)
		}
	}

}
