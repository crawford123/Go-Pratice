package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	var ops unit64
	
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUnit64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
			
		}()
	}
	
	time.Sleep(time.Second)
	
	opsFinal := atomic.LoadUnit64(&ops)
	fmt.Println("ops:", opsFinal)

}