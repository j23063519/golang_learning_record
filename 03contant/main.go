package main

import "fmt"

func main() {
	// 1.
	const TT = "test"
	// TT = "rr"
	// cannot assign to TT (untyped string constant "test")
	fmt.Print("TT:", TT, "\n")

	// 2.
	// iota represents successive integer constants 0, 1, 2
	const (
		A = iota // 0
		B        // 1
		C        // 2
		D        // 3
	)
	fmt.Print("A:", A, "\n")
	fmt.Print("B:", B, "\n")
	fmt.Print("C:", C, "\n")
	fmt.Print("D:", D, "\n")
	// output:
	// A:0
	// B:1
	// C:2
	// D:3

	// 3.
	const (
		A1 = iota + 100 // 100
		B2              // 101
		C3              // 102
		D4              // 103
	)
	fmt.Print("\n\nA1:", A1, "\n")
	fmt.Print("B2:", B2, "\n")
	fmt.Print("C3:", C3, "\n")
	fmt.Print("D4:", D4, "\n")
	// output:
	// A1:100
	// B2:101
	// C3:102
	// D4:103

	// 4.
	const (
		A11 = 1 << iota // 1
		B21             // 2
		C31             // 4
		D41             // 8
	)
	fmt.Print("\n\nA11:", A11, "\n")
	fmt.Print("B21:", B21, "\n")
	fmt.Print("C31:", C31, "\n")
	fmt.Print("D41:", D41, "\n")
	// output:
	// A11:1
	// B21:2
	// C31:4
	// D41:8
}
