// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i = 0

func incrementing() {
    //TODO: increment i 1000000 times
    i += 1
}

func decrementing() {
    //TODO: decrement i 1000000 times
    i -= 1
}

func thread_1(){
    for j := 0; j < 1000000; j++{
        incrementing()
    }
}
func thread_2(){
    for k := 0; k < 1000000; k++{
        decrementing()
    }
}

func main() {
    // What does GOMAXPROCS do? What happens if you set it to 1?
    runtime.GOMAXPROCS(2)    
	
    // TODO: Spawn both functions as goroutines
    //spawn thread_1
    go thread_1()
    //spawn stread_2
	go thread_2()
    // We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
    // We will do it properly with channels soon. For now: Sleep.
    time.Sleep(500*time.Millisecond)
    Println("The magic number is:", i)
}
