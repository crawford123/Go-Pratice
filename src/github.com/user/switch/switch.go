package main

import "fmt"
import "time"

func main(){

	i := 2
	fmt.Println("Write ", i, " as ")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	//Monday
	fmt.Println(time.Now().Weekday())
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
    default:
		fmt.Println("It's a weekday")
	}
	
	t := time.Now()
	//2019-07-29 23:06:46.0560485 +0800 CST m=+0.109111601
	fmt.Println(t)
	//23
	fmt.Println(t.Hour())
	switch{
	case t.Hour() < 12:
		fmt.Println("It's before noon")
    default:
		fmt.Println("It's after noon")
	}
	
	whatAmI := func(i interface{}){
		switch t := i.(type){
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
	    default:
            fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}