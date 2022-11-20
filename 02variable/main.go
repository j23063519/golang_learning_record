package main

import "fmt"

// These variables are global variables
var g = 1
var i = 123
var k = 100

func main() {
	// variable declaration
	// 1.
	var a int = 0
	fmt.Print("a:", a, "\n")
	// output:
	// a:0

	// 2.
	var b int
	b = 1
	fmt.Print("b:", b, "\n")
	// output:
	// b:1

	// 3. Short variable declarations
	c := 2
	fmt.Print("c:", c, "\n")
	// output:
	// c:2

	// note:
	// declare(:=) and assign(=) is different
	// (:=) often used in init multiple variables and for loop and scope part
	// (var =) often used in clear data type and long lifecycle
	var d int = 4
	e, d := 100, 99 // Variable d has been declared so that it just assign new value
	f, d := 5, 6
	fmt.Print(d, e, f, "\n")

	// 4.
	h := 10
	fmt.Print("\n\n&h:", &h, " h:", h, " &g:", &g, " g:", g, "\n")

	g, h := 100, 99
	fmt.Print("&h:", &h, " h:", h, " &g:", &g, " g:", g, "\n")
	// output:
	// &h:0x140000140a0 h:10 &g:0x104b95070 g:global
	// &h:0x140000140a0 h:99 &g:0x140000100a0 g:tt
	// note: If I use := to declare new value that will renew g and its memory position

	// 5.
	fmt.Print("\n\n&i:", &i, " i:", i, "\n")
	i = 122
	fmt.Print("&i:", &i, " i:", i, "\n")
	i := 121
	fmt.Print("&i:", &i, " i:", i, "\n")
	i, j := 120, 120
	fmt.Print("&i:", &i, " i:", i, " &j:", &j, " j:", j, "\n")
	// output:
	// &i:0x10084c390 i:123
	// &i:0x10084c390 i:122
	// &i:0x14000126020 i:121
	// &i:0x14000126020 i:120 &j:0x14000126028 j:120
	// note: If I use (=) assign new value that will not renew g and its memory position instead of (:=)

	// 6.
	fmt.Print("\n\n&k:", &k, " k:", k, "\n")
	var k = 100 // In function renew k variable
	fmt.Print("&k:", &k, " k:", k, "\n")
	// var k, l = 100, 100 // I can't use var anymore because I have already declared k. But I can use (:=) to assign new value.
	k, l := 100, 100
	fmt.Print("&k:", &k, " k:", k, " &l:", &l, " l:", l, "\n")

	if true {
		// In if this area I can use var because it is new area.
		var k, l = 100, 100
		fmt.Print("&k:", &k, " k:", k, " &l:", &l, " l:", l, "\n")
	}

	fmt.Print("&k:", &k, " k:", k, " &l:", &l, " l:", l, "\n")
	// output:
	// &k:0x1028dc398 k:100
	// &k:0x1400011c030 k:100
	// &k:0x1400011c030 k:100 &l:0x1400011c038 l:100
	// &k:0x1400011c040 k:100 &l:0x1400011c048 l:100
	// &k:0x1400011c030 k:100 &l:0x1400011c038 l:100

	// 7.
	// note: If I see this below:
	// connect, err := ...(server)
	// or
	// result, err := ...
	// that means left side variables are scoped variables not the global variables

	// 8.
	// note: If I want to declared global variables I only can use var instead of :=
}
