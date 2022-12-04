package main

import "fmt"

func main() {
	// %d: digit
	// %c: char
	// %s: string
	// %v: value

	var a int = 10
	fmt.Printf("a: %v\n", a)
	// output:
	// a: 10

	var b int = 20
	// fmt.Print("b: %v\n", b)
	// output:
	// b: %v
	// 20%
	fmt.Print("b: ", b, "\n")
	// output:
	// b: 20

	var c int = 30
	fmt.Println("c:", c)
	// output:
	// c: 30

	// advance to next
	// type Test struct {
	// 	A bool
	// 	B int
	// 	C string
	// 	D float64
	// }
	// fmt.Printf("%v \n", Test{})
	// fmt.Printf("%+v \n", Test{})
	// fmt.Printf("%#v \n", Test{})
	// output:
	// {false 0  0}
	// {A:false B:0 C: D:0}
	// main.Test{A:false, B:0, C:"", D:0}

	var d int = 1
	fmt.Printf("d: %v", d)
	// output:
	// d: 1% // % => it means that no new line

	// var e int
	// fmt.Scanf("\n%d", &e)
	// fmt.Println(e)
	// output:
	// what you type

	a1 := "1"
	a2 := "2"
	a3 := "3"

	res1 := fmt.Sprintln("\n", a1, a2, a3)
	fmt.Println(res1)
	// output:
	// 1 2 3
	res2 := fmt.Sprint(a1, a2, a3)
	fmt.Println(res2)
	// output:
	// 123
}
