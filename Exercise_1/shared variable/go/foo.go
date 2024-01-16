// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	//"time"
)

var i = 0

func i_handle(incrChan <-chan bool, decrChan <-chan bool, done <-chan bool) {
	//TODO: increment i 1000000 times
	for {
		select {
		case <-incrChan:
			i += 1
		case <-decrChan:
			i -= 1
		case <-done:
			return
		}
	}
}

func thread_1(incrChan chan<- bool, done chan<- bool) {
	for j := 0; j < 1000000; j++ {
		incrChan <- true
	}
	done <- true
}
func thread_2(decrChan chan<- bool, done chan<- bool) {
	for k := 0; k < 999999; k++ {
		decrChan <- true
	}
	done <- true
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(3)
	// TODO: Spawn both functions as goroutines
	incrChan := make(chan bool)
	decrChan := make(chan bool)
	done := make(chan bool)
	//spawn thread_1
	go i_handle(incrChan, decrChan, done)
	go thread_1(incrChan, done)
	//spawn stread_2
	go thread_2(decrChan, done)
	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	<-done
	<-done
	done <- true
	//time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}
