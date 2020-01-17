package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	resp, err := http.Get("http://localhost:8090/hello")
	//resp, err := http.Get("http://baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response StatusCode:", resp.StatusCode)
	fmt.Println("Response Body:", resp.Body)
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
