package main

import "fmt"

func main(){

    //make(map[key-type]val-type)
	m := make(map[string]int)
	
	m["k1"] = 7
	m["k2"] = 13
	
	fmt.Println("map:", m)
	
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	
	fmt.Println("len: ", len(m))
	
	delete(m, "k2")
	fmt.Println("map: ", m)
	
	_, prs := m["k2"]
    result := m["k2"]
	_, prs1 := m["k1"]
	//output false
	fmt.Println("prs", prs)
	//output true
	fmt.Println("prs1", prs1)
	// output 0
	fmt.Println("result",result)
	n := map[string]int{"foo":1, "bar":2}
	fmt.Println("map", n)
	
	strMap := make(map[string]string)
	strMap["a1"] = "Hi"
	strMap["a2"] = "Hello"
	fmt.Println("strMap",strMap)
	delete(strMap,"a1")
	fmt.Println("newStrMap",strMap)
	_, result1 := strMap["a3"]
	result2 := strMap["a3"]
	//output false
	fmt.Println("result1",result1)
	//output ""
	fmt.Println("result2",result2)
	


}