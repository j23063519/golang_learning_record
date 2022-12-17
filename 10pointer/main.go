package main

import (
	"fmt"
)

func main() {
	// 指標
	// 可以指向 變數 的記憶體位置，或是struct物件本身
	// & 取變數的位址
	// * 取變數的數值

	var ptrInt *int
	fmt.Println("Value of pointer is: ", ptrInt)
	fmt.Printf("Type of ptrInt is: %T\n", ptrInt)
	// output:
	// Value of pointer is:  <nil>
	// Type of ptrInt is: *int

	var ptrString *string
	fmt.Println("Value of pointer is: ", ptrString)
	fmt.Printf("Type of ptrString is: %T\n", ptrString)
	// output:
	// Value of pointer is:  <nil>
	// Type of ptrString is: *string
}
