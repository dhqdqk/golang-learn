package main

import "fmt"
import "errors"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func myarg(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}

func main() {
	//hello world
	fmt.Printf("Hello, world;\n")

	// var
	fmt.Println("******var******")
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
	fmt.Println("******array******")
	var a [3]int
	a[0] = 1
	b := [3]int{4, 5, 6}
	ab := [2][3]int{[3]int{1, 0, 0}, [3]int{4, 5, 6}}
	ba := [2][3]int{b, a}
	fmt.Printf("array a: %v\narray b: %v\narray ab: %v\narray ba:%v\n", a, b, ab, ba)
	//slice
	var s1 []int
	fmt.Printf("sllice s1: %v\n", s1)

	// map
	fmt.Println("******map******")
	mdict := make(map[string]int)
	mdict["one"] = 1
	mdict["two"] = 2
	mdict["three"] = 3
	fmt.Println("mdict[`three`] is ", mdict["three"])

	// for
	fmt.Println("******for******")
	for k, v := range mdict {
		fmt.Printf("%v: %v\n", k, v)
	}

	// switch
	fmt.Println("******switch******")
	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
	default:
		fmt.Println("default case")
	}

	// func
	fmt.Println("******func******")
	x := 3
	y := 4
	z := 5
	max_xy := max(x, y)
	max_xz := max(x, z)

	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z))

	// func with arg
	fmt.Println("******func-with-arg******")
	myarg(1, 2, 3, 4, 5)

	// defer
	fmt.Println("******defer******")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}

}
