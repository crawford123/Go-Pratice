package main

import "fmt"
import "math"

func main() {
	//fmt输出格式 \n\t\r
	fmt.Println("hello world")
	fmt.Printf("%t\n", 1 == 2)
	fmt.Printf("二进制：%b\n", 255)
	fmt.Printf("十六进制：%X\n", 255)
	fmt.Printf("十进制：%d\n", 255)
	fmt.Printf("浮点数：%f\n", math.Pi)
	fmt.Printf("字符串：%s\n", "hello world")

	//变量和常量
	//go是静态类型的语言
	//声明初始化一个变量
	//var x int = 100
	//var str string = "hello world"
	//声明初始化多个变量
	//var i, j, k int = 1, 2, 3
	//不用指明类型，通过初始化值来推导
	//var b = true
	//var c = "123"
	//另一种定义变量的方式
	//y := 100 // 等价于 var x int = 100
	//常量很简单，使用const关键字
	const s string = "hello world"
	const pi float32 = 3.1415926

	//数组
	var a [5]int
	fmt.Println("array a:", a)

	a[1] = 10
	a[3] = 30
	fmt.Println("assign:", a)
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("init:", b)

	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("2d: ", c)

	//数组的切片操作
	d := [5]int{1, 2, 3, 4, 5}
	fmt.Println("d: ", d)
	e := d[2:4] // // a[2] 和 a[3]，但不包括a[4]
	fmt.Println("e: ", e)
	f := d[:4] // 从 a[0]到a[4]，但不包括a[4]
	fmt.Println("f: ", f)
	g := d[2:] //从 a[2]到a[4]，且包括a[2]
	fmt.Println("g: ", g)

	//分支循环语句
	//if语句 if 语句没有圆括号，而必需要有花括号
	var x int = 100
	if x%2 == 0 {
		fmt.Println("这是偶数！")
	}
	//if - else
	if x%2 == 0 {
		fmt.Println("这是偶数！")
	} else {
		fmt.Println("这是奇数！")
	}
	//多分支
	if x < 0 {
		fmt.Println("这是负数！")
	} else if x == 0 {
		fmt.Println("这是0！")
	} else {
		fmt.Println("这是正数！")
	}
	//switch 语句
	//switch语句没有break，还可以使用逗号case多个值
	var i int = 0
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5, 6:
		fmt.Println("four,five,six")
	default:
		fmt.Println("invalid value!")
	}

	//for 语句
	//下面再来看看for的三种形式：（注意：Go语言中没有while）
	//经典的for语句 init; condition; post
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("========")
	//精简的for语句 condition
	k := 1
	for k < 10 {
		fmt.Println(k)
		k++
	}
	fmt.Println("========")
	//死循环的for语句 相当于for(;;)
	j := 1
	for {
		if j > 10 {
			break
		}
		fmt.Println(j)
		j++
	}

}
