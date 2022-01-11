package main

import "fmt"

func main(){
	// TODO : create channel owner, which creates channel.
	// return receive only channel to caller
	// spine goroutine, which
	// writes data into channel and
	// closes the channel when done.

	owner := func() <-chan int {
		ch := make(chan int)

		go func() {
			defer close(ch)
			for i:=0; i<5; i++ {
				ch <- i
			}
		}()
		return ch
	}

	consumer := func(ch <- chan int) {
		//read values from channel
		for v := range ch {
			fmt.Printf("Received: %d\n", v)
		}
		fmt.Println("done receiving!")
	}

	ch := owner()
	consumer(ch)
}