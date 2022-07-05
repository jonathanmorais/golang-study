package main

import "fmt"

// Prints a greeting message using values received in
// the channel
func first(a chan string) {

	first := <-a // receiving value from channel
	fmt.Println(first)
}

func second(b chan string) {
	second := <-b
	fmt.Println(second)
}

func third(c chan string) {
	third := <-c
	fmt.Println(third)
}

func main() {

	// Making a channel of value type string
	a := make(chan string)
	b := make(chan string)
	c := make(chan string)

	// Starting a concurrent goroutine
	go first(a)
	go second(b)
	go third(c)

	// Sending values to the channel c
	c <- "make"
	a <- "be"
	b <- "yourself"

	// Closing channel
	close(c)
	close(a)
	close(b)
}
