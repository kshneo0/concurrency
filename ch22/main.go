package main

import "fmt"

// TODO: Build a Pipeline
// generator () -> square() -> print

// generator - convertes a list of integers to a channel

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out

}


func main() {
	// set up the pipeline

	// run the last stage of pipeline
	// recevie the values from square stage
	//print each ont, until channel is closed.

	// ch := generator(2,3)
	// out := square(ch)
	// for n := range out {
	// 	fmt.Println(n)
	// }

	for n := range square(generator(2,3)) {
		fmt.Println(n)
	}

}