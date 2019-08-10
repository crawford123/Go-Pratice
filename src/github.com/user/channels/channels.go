package main

import "fmt"


/*func test(){
		test := <- messages
		fmt.Println("test的值为：",test)
	}*/
	
func main(){

//Create a new channel with make(chan val-type)
	messages := make(chan string)
	
	
	go func(){
	//Send a value into a channel using the channel <- syntax
		messages <- "ping"
	}()
	
	//go test()
	
	//The <-channel syntax receives a value from the channel
	msg := <- messages
	fmt.Println(msg)
	
}