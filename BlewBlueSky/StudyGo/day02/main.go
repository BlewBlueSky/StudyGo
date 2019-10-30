package main

import (
	"unicode"
	"fmt"
)

// 全局变量m
var m = 100
var name string = "Q1mi"
var age int = 18
var sp string = "hello沙河"
// 常变量
const (
	pi = 3.1415
	e  = 2.7182
)

func main() {
	n := 10
	m := 200 // 此处声明局部变量m
	
	
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(m, n, name, age, pi, e)
	fmt.Printf(`%T`, n)				//打印数据类型
	fmt.Println()
	for _, r := range sp { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	//寻找汉字
	s1 := "我是汉字abcd"
    var count int
    for _, v := range s1 {
        if unicode.Is(unicode.Han, v) {
            count++
        }
	}
	fmt.Printf("\n%s\n",s1)
	fmt.Printf("找到汉字：%d\n",count)
	//匿名变量
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)

}
func foo() (int, string) {
	return 10, "Q1mi"
}
