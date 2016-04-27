package main

import "fmt"
import "errors"

func main() {
	//hello world
	fmt.Printf("Hello, world;\n")
	// var
	var c complex64 = 5 + 5i
	fmt.Printf("Value is: %v\n", c)
	// string
	s := "hello"
	s = "d" + s[1:]
	fmt.Printf("%s\n", s)
	// error
	err := errors.New("Error:to duanhq: dhqdqk\n")
	if err != nil {
		fmt.Print(err)
	}
	// array
	var a [3]int
	a[0] = 1
	b := [3]int{4, 5, 6}
	ab := [2][3]int{[3]int{1, 0, 0}, [3]int{4, 5, 6}}
	ba := [2][3]int{b, a}
	fmt.Printf("array a: %v\narray b: %v\narray ab: %v\narray ba:%v\n", a, b, ab, ba)
	//slice
	var s1 []int
	fmt.Printf("sllice s1: %v\n", s1)

}
