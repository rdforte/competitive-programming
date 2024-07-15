package main

import (
	"fmt"
	"time"
)

// Senders should be responsible for closing their own channels.
// Sending on a closed channel causes a panic
// Not closing your channels can cause a deadlock
// Closed channels are non-blocking so anything reading from that channel will read its zero value ie: int = 0
// Nil channels are blocking. This is how we avoid are select statement from spinning and constantly printing 0's.

func main() {
	evenChan := make(chan int)
	oddChan := make(chan int)

	combinedChan := make(chan string)

	go func() {
		for evenChan != nil || oddChan != nil {
			time.Sleep(time.Second * 1)
			select {
			case v, open := <-evenChan:
				if !open {
					evenChan = nil
				} else {
					combinedChan <- fmt.Sprintf("received evenChan: %d", v)
				}
			case v, open := <-oddChan:
				if !open {
					oddChan = nil
				} else {
					combinedChan <- fmt.Sprintf("received oddChan: %d", v)
				}
			}
		}
		close(combinedChan)
	}()

	go func() {
		for i := 1; i < 10; i++ {
			// close chan if i is 6
			if i == 6 {
				close(evenChan)
			}
			// post to even chan if i is even
			if i < 6 && i%2 == 0 {
				evenChan <- i
			}
			// close odd chan if i is 9
			if i == 9 {
				close(oddChan)
			}
			// post to odd chan if i is odd
			if i < 9 && i%2 == 1 {
				oddChan <- i
			}
		}
	}()

	for s := range combinedChan {
		fmt.Println(s)
	}

	fmt.Println("DONE")
}
