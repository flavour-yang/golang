package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variableZeroValue() {
	var a string
	var b string
	a = "aaa"
	b = "bbb"
	fmt.Printf("%q %q", a, b)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c)) //

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1) // 欧拉公式
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)         // 欧拉公式
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)  // 欧拉公式
}

// 类型强制转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64((a*a + b*b))))
	fmt.Println(c)
}

// go语言的常量
func const1() {
	const filename = "abc.txt"
	const a, b = 3, 5
	c := int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// go的枚举类型: go没有枚举类型关键字
func enums() {
	// const (
	// 	a = 0
	// 	b = 1
	// 	c = 2
	// )
	// const (
	// 	a = 2 * iota
	// 	b
	// 	c
	// )
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(b, kb, mb, gb, tb)
}

func main() {
	fmt.Println("Hello Word")
	variableZeroValue()
	euler()
	triangle()
	const1()
	enums()
}

// 基本变量类型
// bool,string,
// (u)int,(u)int8,(u)int16,(u)int32,(u)int64, 指针: uintptr
// byte, rune
// float32,float64,
// 虚数 comlpex64,complex128
