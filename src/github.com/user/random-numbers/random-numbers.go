package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// rand.Intn returns a random int n, 0 <= n < 100.
	//固定的值
	fmt.Print(rand.Intn(100), ",")
	fmt.Println(rand.Intn(100))
	fmt.Println()

	//Float64 returns a float64 f, 0.0 <= f < 1.0.
	//固定的值
	fmt.Println(rand.Float64())

	// 5.0 <= f' < 10.0.
	//随机变动的值
	fmt.Print((rand.Float64() * 5) + 5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	fmt.Println("unixNano:", time.Now().UnixNano())
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	//随机变动的值
	fmt.Print(r1.Intn(100), ",")
	fmt.Println(r1.Intn(100))
	fmt.Println()

	//5,87
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Println(r2.Intn(100))
	fmt.Println()
	//5,87
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Println(r3.Intn(100))

}
