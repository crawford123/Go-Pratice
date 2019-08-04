package main

import "fmt"

type person struct {
	name string
	age int
}

func main(){
	
	fmt.Println(person{"Bob",20})
	
	fmt.Println(person{name:"Alice",age :30})
	
	fmt.Println(person{name: "Fred"})
	
	fmt.Println(person{age:20})
		
	fmt.Println(&person{name: "Ann",age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	
	s1 := s
	fmt.Println(s1)
	sp := &s
	fmt.Println(sp)
	sp1 := &person{name: "Tony",age: 24}
	fmt.Println(sp.age)
	fmt.Println(s)
	fmt.Println(sp1)
	
	sp.age = 51
	fmt.Println(sp.age)


}