package main

import "fmt"

func main() {

	 queue := make(chan string, 2)
	//test := make(chan string)
	
	//test <- "121"
	//fmt.Println(<-test)
	//close(test)
	
	queue <- "one"
	queue <- "two"
	close(queue)
	
	for elem := range queue {
		fmt.Println(elem)
	}
	
}