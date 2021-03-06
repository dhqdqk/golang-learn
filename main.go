package main

import "fmt"
import "errors"
import "reflect"
import "runtime"

// goto
func mygoto() {
	j := 0
Here: //goto-tag
	fmt.Println(j)
	j++
	if j > 5 {
		goto Out
	}
	goto Here
Out:
	fmt.Println("out of goto")
}

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

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box // a list of boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func gosum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
		fmt.Println(v)
	}
	c <- total // send total to c
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
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

	// iota
	fmt.Println("*******iota******")

	fmt.Printf("WHITE: %v\n", WHITE)
	fmt.Printf("BLUE: %v\n", BLUE)

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

	// map[keyType]valueType
	fmt.Println("******map******")
	mdict := make(map[string]int)
	mdict["one"] = 1
	mdict["two"] = 2
	mdict["three"] = 3
	fmt.Println("mdict[`three`] is ", mdict["three"])

	// if-else
	integer := 3
	if integer == 3 {
		fmt.Println("integer == 3\n")
	} else if integer < 3 {
		fmt.Println("integer is not equal to 3\n")
	}

	// for
	fmt.Println("******for******")
	for k, v := range mdict {
		fmt.Printf("%v: %v\n", k, v)
	}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum is equal to ", sum)

	sum1 := 1
	for sum1 < 500 {
		sum1 += sum1
	}
	fmt.Println("now sum is equal to ", sum1)

	// switch
	fmt.Println("******switch******")
	integers := 6
	switch integers {
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

	// goto
	mygoto()

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

	// struct
	fmt.Println("*******struct*******")
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Printf("The volume of the first one is : %v (%v, %v, %v)\n", boxes[0].Volume(), boxes[0].width, boxes[0].height, boxes[0].depth)
	fmt.Println("The bigest one is", boxes.BiggestColor().String())
	fmt.Println("The bigest one is :", boxes.BiggestColor())

	// reflect
	fmt.Println("********reflect*******")
	var re float64 = 3.14
	re_t := reflect.TypeOf(re)
	re_v := reflect.ValueOf(re)
	fmt.Println("the type of re is :", re_t, re_v)

	// runtime-goroutine
	fmt.Println("********runtime & goroutine********")
	go say("world") // create a new goroutine
	say("hello")    // current goroutine run at time
	go_a := []int{7, 2, 8, -9, 4, 0}
	go_c := make(chan int)
	fmt.Println(go_a[:len(go_a)/2])
	go gosum(go_a[:len(go_a)/2], go_c)
	go gosum(go_a[len(go_a)/2:], go_c)
	go_x, go_y := <-go_c, <-go_c // receive from c
	fmt.Println(go_x, go_y, go_x+go_y)

	// channel
	ch_c := make(chan int, 10)
	go fibonacci(10, ch_c)
	for i := range ch_c {
		fmt.Println(i)
	}
}
