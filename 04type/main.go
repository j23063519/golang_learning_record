package main

import "fmt"

// globalDeclared := 1 => get error

func main() {
	// bool is the set of boolean values, true and false.
	var a bool = false
	// a = 0 => get error
	// a = "a" => get error
	fmt.Println(a)
	// output:
	// false

	// advance to next
	var truth = 3 <= 5
	var falsehood = 10 != 10
	var res1 = 10 > 20 && 5 == 5     // Second operand is not evaluated since first evaluates to false
	var res2 = 2*2 == 4 || 10%3 == 0 // Second operand is not evaluated since first evaluates to true
	fmt.Println(truth, falsehood, res1, res2)
	// output:
	// true,false,false,true

	// int8: int8 is the set of all signed 8-bit integers. Range: -128 through 127. 2^8
	// int16: int16 is the set of all signed 16-bit integers. Range: -32768 through 32767. 2^16
	// int32: int32 is the set of all signed 32-bit integers. Range: -2147483648 through 2147483647. 2^32
	// int64: int64 is the set of all signed 64-bit integers. Range: -9223372036854775808 through 9223372036854775807. 2^64
	// int is a signed integer type that is at least 32 bits in size. It is a distinct type, however, and not an alias for, say, int32.
	// int 型別的大小取決於平台。在 32 位元系統上，它的大小為 32 位元，在 64 位元系統上，它的大小為 64 位元。
	var b int = 1
	// b = "1" => get error
	// b = 10000000000000000000 => get error because over limits
	fmt.Println(b)
	// output:
	// 1

	// uint8: uint8 is the set of all unsigned 8-bit integers. Range: 0 through 255. 2^8
	// uint16: uint16 is the set of all unsigned 16-bit integers. Range: 0 through 65535. 2^16
	// uint32: uint32 is the set of all unsigned 32-bit integers. Range: 0 through 4294967295. 2^32
	// uint64: uint64 is the set of all unsigned 64-bit integers. Range: 0 through 18446744073709551615. 2^64
	// uint is an unsigned integer type that is at least 32 bits in size. It is a distinct type, however, and not an alias for, say, uint32.
	// uint 型別的大小取決於平台。在 32 位元系統上，它的大小為 32 位元，在 64 位元系統上，它的大小為 64 位元。
	var c uint = 10
	// c = -1 get error
	// c = 100000000000000000000 => get error because over limits
	fmt.Println(c)
	// output:
	// 10

	// advance to next
	// 在 Golang 中，你可以使用前綴 0 來宣告八進制數字，以及使用前綴 0x 或 0X 來宣告十六進制數字。
	var myHexNumber = 0xFF  // Use prefix '0x' or '0X' for declaring hexadecimal numbers
	var myOctalNumber = 034 // Use prefix '0' for declaring octal numbers
	fmt.Printf("%#x, %#o\n", myHexNumber, myOctalNumber)
	// output:
	// 0xff, 034
	var initInt uint
	fmt.Println(initInt)
	// output:
	// 0

	// float64: float64 is the set of all IEEE-754 64-bit floating-point numbers.
	// float32 is the set of all IEEE-754 32-bit floating-point numbers.
	var d float32 = 10.01
	// d = "2" get error
	fmt.Println(d)
	// output:
	// 10.01

	// complex128: complex128 is the set of all complex numbers with float64 real and imaginary parts.
	// complex64 is the set of all complex numbers with float32 real and imaginary parts.
	var e complex128 = 2i + 1.9
	fmt.Println(e)
	// output:
	// (1.9+2i)

	// string is the set of all strings of 8-bit bytes, conventionally but not necessarily representing UTF-8-encoded text. A string may be empty, but not nil. Values of string type are immutable.
	var f string = "hello"
	fmt.Println(f)
	// output:
	// hello

	// advance to next
	var bio string = `test
	test`
	fmt.Println(bio)
	// output:
	// test
	//         test

	// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is used, by convention, to distinguish byte values from 8-bit unsigned integer values.
	var g byte = 'a'
	fmt.Println(g)
	// output:
	// 97
	// other example
	var h = []byte("Hello, world!")
	fmt.Println(h)
	// output:
	// [72 101 108 108 111 44 32 119 111 114 108 100 33]

	// rune is an alias for int32 and is equivalent to int32 in all ways. It is used, by convention, to distinguish character values from integer values.
	var i rune = '♥' // equals to var i = '♥'
	fmt.Println(i)
	// output:
	// 9829

	// Golang 沒有 char 資料型別。它使用 byte 和 rune 來表示字元值。byte 資料型別用於表示 ASCII 字元，以及 rune 資料型別用於表示以 UTF-8 格式編碼的一組更廣泛的 Unicode 字元。
	// other examples
	// byte 和 rune 資料型別本質上都是整數。舉例來說，具有值 'a' 的 byte 變數會轉換為整數 97。
	var myByte byte = 'a'
	// 具有 Unicode 值 '♥' 的 rune 變數會轉換為相應的 Unicode 代碼點 U + 2665，其中 U+ 代表 Unicode，且數字是十六進制，本質上是整數。
	var myRune rune = '♥'
	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)
	// output:
	// a = 97 and ♥ = U+2665

	// Data type conversion
	var j int64 = 4
	var k int = int(j)
	fmt.Println(k)
	// output:
	// 4
	var l float64 = 2.35
	var res = float64(k) + l
	fmt.Println(res)
	// output:
	// 6.35
	var myInt int = 65
	var myUint uint = uint(myInt)
	var myFloat64 float64 = float64(myUint)
	fmt.Println(myFloat64)
	// output:
	// 65

	// var j int64 = 1
	// var k int = j => get error, cannot use j (variable of type int64) as int value in variable declaration
	// var k int = 1
	// var plus = j + k => invalid operation: j + k (mismatched types int64 and int)
}
