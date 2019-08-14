package main

import "time"
import "fmt"

func main(){
   
	 ticker := time.NewTicker(500 * time.Millisecond)
	 //output: &{0xc0000a0000 {5753664 0 826539961541200 500000000 0x467d30 0xc0000a0000 0}}
	 fmt.Println("ticker:",ticker)
	 //output: 0xc0000a0000
	 fmt.Println("ticker.C:",ticker.C)
     go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")

}