package pratice

import (
	"fmt"
	"time"
)

func sayHello1(done chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, Goroutine!")
		time.Sleep(500 * time.Millisecond)
	}
	done <- true // Sending a signal to the channel to indicate completion of work
}

func main() {
	done := make(chan bool) // Creating a channel for signals

	go sayHello1(done) // Launching sayHello in a goroutine with a channel as an argument

	<-done // Waiting for a signal from the channel. This operation blocks execution until the signal is received.
	fmt.Println("Main function finished")
}
