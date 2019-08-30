package main

import "os"
//import "fmt"


func main(){

	panic("a problem")
	
	_, err := os.Create("/tmp/file")
    //fmt.Println("值为：", os.Create("/tmp/file"))
	if err != nil {
		panic(err)
	}
}