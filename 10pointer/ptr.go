package main

import "fmt"

func ptr() {
	// 簡介
	// 指標是一個變數，用來儲存另一個變數的記憶體位址。
	// 當我們在寫程式時，我們需要在記憶體中儲存一些資料，而資料會儲存在記憶體中的特定位址。記憶體位址會像 0xAFFFF(這是以十六進制來表示記憶體位址)
	// 但是要記住所有記憶體位址並使用是非常困難的，所以才出現了變數的概念，變數只是為儲存資料的記憶體位置所取的名稱
	// 指標也是一個變數。但是他是一種特殊的變數，它是另一個變數的記憶體位址
	//        變數              位址          記憶體
	//                        0x0000
	//     var a int = 2      0x0001           2   <--|
	//                        0x0002                  |
	//                        0x0003                  | 指
	//                        0x0004                  | 向
	//     var p *int = &a    0x0005         0x0001 --|
	//                        0x0006
	//                        0x0007
	// 在上面的例子中，指標 p 含有數值 0x0001，它是變數 a 的位址

	// 宣告指標
	// 使用以下語法宣告一個型別為 int 的指標:
	// 且該指標只能儲存 int 變數的記憶體位址
	// 這時的指標值是nil，由此可知任何未初始化的指標都會是nil
	var p *int
	fmt.Printf("p pointer of type %T\n", p)
	fmt.Printf("p pointer of value %v\n", p)
	// output:
	// p pointer of type *int
	// p pointer of value <nil>

	// 初始化指標
	// 你可以使用另一個變數的記憶體位址來初始化指標。
	// 可以使用 & 運算子取得變數的記憶體位址
	q := 10
	p = &q
	fmt.Printf("Value of q is %v\n", q)
	fmt.Printf("Address of q is %v\n", &q)
	fmt.Printf("Value of pointer p is %v\n", p)
	fmt.Printf("Value of pointer *p is %v\n", *p)
	fmt.Printf("VAddressalue of pointer p is %v\n", &p)
	// output:
	// Value of q is 10
	// Address of q is 0x14000126010
	// Value of pointer p is 0x14000126010
	// Value of pointer *p is 10
	// VAddressalue of pointer p is 0x14000120040
	// 由此可知:
	// 我們用 & 運算子與變數 q 來取得它的記憶體位址，並將記憶體位址指派給指標 p
	// p = &q
	// 若想合併把宣告及賦值同時一起做可以:
	// 方式一
	// r := 20
	// s := &r
	// 方式二
	// var r = 20
	// var s = &r

	// 解參考指標
	// 若要取出指標所指向變數裡的數值，可以用 * 運算子取出。
	// 這個動作稱為 解參考 或 間接取值
	r := 20
	s := &r
	fmt.Println("r = ", r)
	fmt.Println("&r = ", &r)
	fmt.Println("s = ", s)
	fmt.Println("*s = ", *s)
	// output:
	// r =  20
	// &r =  0x1400011c018
	// s =  0x1400011c018
	// *s =  20

	// 除了用*運算子讀取指向變數的值之外，也可以對其進行賦值
	// 賦值之後，就會發現變數r的值會改變
	*s = 10
	fmt.Println("r = ", r)
	fmt.Println("&r = ", &r)
	fmt.Println("s = ", s)
	fmt.Println("*s = ", *s)
	// r =  10
	// &r =  0x1400011c018
	// s =  0x1400011c018
	// *s =  10

	// 使用內建的 new()函數建立指標，此函數會將型別當作參數，並分配足夠的記憶體來容納該行別的值，最後回傳指向該型別的指標
	t := new(int) // Pointer to an integer type
	*t = 100
	fmt.Println("t = ", t)
	fmt.Println("*t = ", *t)
	fmt.Println("&t = ", &t)
	// output:
	// t =  0x1400011c020
	// *t =  100
	// &t =  0x14000114048

	// 指標的指標
	// 指標可以指向任何型別的變數，當然也可以指向另一個指標
	var u = 1.1
	v := &u
	w := &v
	fmt.Println("u = ", u)
	fmt.Println("Address of u = ", &u)
	fmt.Println("v = ", v)
	fmt.Println("address of v = ", &v)
	fmt.Println("w = ", w)
	fmt.Println("address of w = ", &w)
	// Dereferencing a pointer to pointer
	fmt.Println("*w = ", *w)
	fmt.Println("**w = ", **w)
	// output:
	// u =  1.1
	// Address of u =  0x1400011c028
	// v =  0x1400011c028
	// address of v =  0x14000114050
	// w =  0x14000114050
	// address of w =  0x14000114058
	// *w =  0x1400011c028
	// **w =  1.1

	// Go 沒有指標運算
	// Go 不支援對指標進行此類算術運算，但是c/c++可以
	// var x = 15
	// var y = &x
	// var z = y + 1
	// cannot convert 1 (untyped int constant) to *int

	// 雖然go不支持算數運算，但是可以使用 == 運算子比較兩個相同的型別的指標是否相等
	var i = 1
	var j = &i
	var k = &i
	if j == k {
		fmt.Println("Both pointers j and k point to the same variable")
	}
	// output:
	// Both pointers j and k point to the same variable
}
