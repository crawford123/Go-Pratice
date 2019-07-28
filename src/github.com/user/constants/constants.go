package main

import "fmt"
import "math"

const s string = "constant"

func main(){
	fmt.Println(s)
	
	const n = 500000000
	const d = 3e20 / n 
	fmt.Println(d)
	fmt.Println(int64(d))
	
	fmt.Println(math.Sin(n))
	//Sin expects a float64.下面得到的值与上面的相同
	fmt.Println(float64(math.Sin(n)))
}