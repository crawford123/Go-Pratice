package main

import "fmt"

func main() {

//make a channel of strings buffering up to 2 values
	messages := make(chan string,2)
	

    messages  <- "buffered"
    messages <- "channel"
	
	//a:= <-messages
	//b := <-messages
	
	//fmt.Println(a)
	//fmt.Println(b)
	
	// only  fetch once,otherwise occur error
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}