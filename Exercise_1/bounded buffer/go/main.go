
package main

import "fmt"
import "time"


func producer(bufferedChan chan<- int){

    for i := 0; i < 10; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("[producer]: pushing %d\n", i)
        // TODO: push real value to buffer
        bufferedChan <- i
    }

}

func consumer(bufferedChan <-chan int){

    time.Sleep(1 * time.Second)
    for {
        i := <- bufferedChan //TODO: get real value from buffer
        fmt.Printf("[consumer]: %d\n", i)
        time.Sleep(50 * time.Millisecond)
    }
    
}


func main(){
    
    // TODO: make a bounded buffer
    bufferedChan := make(chan int, 5)
    go consumer(bufferedChan)
    go producer(bufferedChan)
    
    select {}
}