package main

import "fmt"

func detail() {
	// slice 是陣列的一部份。slice 建構在陣列之上，但是比陣列更方便以及更靈活的運用。
	// slice 是可以索引的且可以動態調整長度。

	// 宣告 slice =================================================================
	// Slice of type `int`
	var s []int
	fmt.Printf("s: %+v\n", s)
	// Slice 的宣告就像陣列一樣，只不過我們在中括號 [] 中未指定任何長度。
	// output:
	// s: []

	// 建立和初始化 slice
	// 1.使用 slice定數 建立slice
	var a = []int{1, 2, 3, 4, 5}
	fmt.Printf("a: %+v\n", a)
	b := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("b: %+v\n", b)
	// 上述兩個範例 的陳述式右側是 slice定數，它會先建立一個陣列，並返回參考它的slice
	// output:
	// a: [1 2 3 4 5]
	// b: [a b c d e]

	// 2.從陣列建立slice
	// 由於slice是陣列的一部份，因此我們可以從陣列建立slice
	// 若要從陣列建立切片，必須要給予上下限，我們可以指定兩個由冒號分隔的索引low(下限)和high(上限)
	var c = [5]string{"a", "b", "c", "d", "e"}
	var d []string = c[1:4]
	fmt.Printf("c: %+v\n", c)
	fmt.Printf("d: %+v\n", d)
	// output:
	// c: [a b c d e]
	// d: [b c d]
	// slice中的上下限不是必填的，下限的預設值為0，上限則為該切片長度 (上限則為該切片長度這句話，我覺得怪怪的，為什麼不是參考的陣列長度)
	e := c[:]
	fmt.Printf("e: %+v\n", e)
	// output:
	// e: [a b c d e]

	// 3.從一個slice建立另一個slice
	cities := []string{"New York", "London", "Chicago", "Beijing", "Delhi", "Mumbai", "Bangalore", "Hyderabad", "Hong Kong"}
	asianCities := cities[3:]
	indianCities := asianCities[1:5]
	fmt.Printf("cities: %+v\n", cities)
	fmt.Printf("asianCities: %+v\n", asianCities)
	fmt.Printf("indianCities: %+v\n", indianCities)
	// output:
	// cities: [New York London Chicago Beijing Delhi Mumbai Bangalore Hyderabad Hong Kong]
	// asianCities: [Beijing Delhi Mumbai Bangalore Hyderabad Hong Kong]
	// indianCities: [Delhi Mumbai Bangalore Hyderabad]

	// 修改slice =================================================================
	// slice是參考型別，它會參考底層陣列。若修改slice的元素，將會修改參考陣列中的相應元素，而參考相同底層陣列的slice的內部元素也會跟著變換
	f := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	g := f[1:]
	h := f[3:]
	fmt.Println("------- Before Modifications -------")
	fmt.Printf("f: %+v\n", f)
	fmt.Printf("g: %+v\n", g)
	fmt.Printf("h: %+v\n", h)
	g[0] = "TUE"
	g[1] = "WED"
	g[2] = "THU"
	h[1] = "FRIDAY"
	fmt.Println("------- After Modifications -------")
	fmt.Printf("f: %+v\n", f)
	fmt.Printf("g: %+v\n", g)
	fmt.Printf("h: %+v\n", h)
	// output:
	// ------- Before Modifications -------
	// f: [Mon Tue Wed Thu Fri Sat Sun]
	// g: [Tue Wed Thu Fri Sat Sun]
	// h: [Thu Fri Sat Sun]
	// ------- After Modifications -------
	// f: [Mon TUE WED THU FRIDAY Sat Sun]
	// g: [TUE WED THU FRIDAY Sat Sun]
	// h: [THU FRIDAY Sat Sun]

	// slice 的長度和容量 =================================================================
	// slice 由三樣東西組成:
	// 指標:指向底層的陣列
	// 長度:slice 包含陣列的長度
	// 容量:能夠容納多少的元素(該部分可以增加到的最大大小)
	// 讓我們用以下陣列和從中得到的 Slice 作為範例：
	var i = [6]int{10, 20, 30, 40, 50, 60}
	var j = i[1:4]
	fmt.Printf("j = %v, len = %d, cap = %d\n", j, len(j), cap(j))
	// output:
	// j = [20 30 40], len = 3, cap = 5

	// 透過重新劃分，可以將 Slice 的長度擴充到其容量。任何嘗試將其長度擴充到可用容量之外的行為將導致執行時期錯誤。
	// 查看以下範例來了解如何重新劃分指定 Slice 來更改其長度和容量：
	k := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("\nOriginal Slice")
	fmt.Printf("k = %v, len = %d, cap = %d\n", k, len(k), cap(k))
	k = k[1:5]
	fmt.Println("\nAfter slicing from index 1 to 5")
	fmt.Printf("k = %v, len = %d, cap = %d\n", k, len(k), cap(k))
	k = k[:8]
	fmt.Println("\nAfter extending the length")
	fmt.Printf("k = %v, len = %d, cap = %d\n", k, len(k), cap(k))
	k = k[2:]
	fmt.Println("\nAfter dropping the first two elements")
	fmt.Printf("k = %v, len = %d, cap = %d\n", k, len(k), cap(k))
	// output:
	// Original Slice
	// k = [10 20 30 40 50 60 70 80 90 100], len = 10, cap = 10

	// After slicing from index 1 to 5
	// k = [20 30 40 50], len = 4, cap = 9

	// After extending the length
	// k = [20 30 40 50 60 70 80 90], len = 8, cap = 9

	// After dropping the first two elements
	// k = [40 50 60 70 80 90], len = 6, cap = 7

	// 使用內建的 make函數建立slice =================================================================
	// make 函數需要型別、長度和容量(選填)。它分配了長度與給定容量相同的底層陣列，
	// 並返回參考陣列的slice
	l := make([]int, 5, 10)
	fmt.Printf("l = %v, len = %d, cap = %d\n", l, len(l), cap(l))
	// output:
	// l = [0 0 0 0 0], len = 5, cap = 10

	// make() 函數中的容量參數是非必要的。如果省略，則預設為指定的長度：
	m := make([]int, 5)
	fmt.Printf("m = %v, len = %d, cap = %d\n", m, len(m), cap(m))
	// output:
	// m = [0 0 0 0 0], len = 5, cap = 5

	// slice 的零值 =================================================================
	// slice 的零值是 nil。具有零值的 Slice 沒有任何底層陣列，且長度和容量為 0：
	var n []int
	fmt.Printf("n = %v, len = %d, cap = %d\n", n, len(n), cap(n))
	if n == nil {
		fmt.Println("n is nil")
	}
	// output:
	// n = [], len = 0, cap = 0
	// n is nil

	// slice 函數 =================================================================
	// 1.copy() 函數將元素從一個 Slice 複製到另一個 Slice
	// func copy(dst, src []T) int
	// 它需要兩個 Slice：來源 Slice 和目的 Slice。然後它會將元素從來源複製到目的，並返回複製的元素數量。
	// 複製的元素數量將是 len(src) 和 len(dst) 的最小值。
	src := []string{"Sublime", "VSCode", "IntelliJ", "Eclipse"}
	dest := make([]string, 2)
	numElementsCopied := copy(dest, src)
	fmt.Println("src = ", src)
	fmt.Println("dest = ", dest)
	fmt.Println("Number of elements copied from src to dest = ", numElementsCopied)
	// output:
	// src =  [Sublime VSCode IntelliJ Eclipse]
	// dest =  [Sublime VSCode]
	// Number of elements copied from src to dest =  2

	// 2.append() 函數:在給定的 Slice 結尾附加新元素
	// func append(s []T, x ...T) []T
	// 如果給定的 Slice 沒有足夠的容量來容納新元素，則將分配具有更大容量的新底層陣列。
	// 現有 Slice 的底層陣列中的所有元素都將複製到新陣列，然後附加新元素。
	// 但是，如果 Slice 具有足夠的容量來容納新元素，則 append() 函數將再使用其底層陣列並將新元素附加到同一陣列中。

	// 容量不夠：
	notBigSlice1 := []string{"C", "C++", "Java"}
	notBigSlice2 := append(notBigSlice1, "Python", "Ruby", "Go")
	fmt.Printf("notBigSlice1 = %v, len = %d, cap = %d\n", notBigSlice1, len(notBigSlice1), cap(notBigSlice1))
	fmt.Printf("notBigSlice2 = %v, len = %d, cap = %d\n", notBigSlice2, len(notBigSlice2), cap(notBigSlice2))
	notBigSlice1[0] = "C#"
	fmt.Println("\nnotBigSlice1 = ", notBigSlice1)
	fmt.Println("notBigSlice2 = ", notBigSlice2)
	// output:
	// notBigSlice1 = [C C++ Java], len = 3, cap = 3
	// notBigSlice2 = [C C++ Java Python Ruby Go], len = 6, cap = 6
	// notBigSlice1 =  [C# C++ Java]
	// notBigSlice2 =  [C C++ Java Python Ruby Go]

	// 容量足夠：
	bigSlice1 := make([]string, 3, 10)
	copy(bigSlice1, []string{"C", "C++", "Java"})
	bigSlice2 := append(bigSlice1, "Python", "Ruby", "Go")
	fmt.Printf("bigSlice1 = %v, len = %d, cap = %d\n", bigSlice1, len(bigSlice1), cap(bigSlice1))
	fmt.Printf("bigSlice2 = %v, len = %d, cap = %d\n", bigSlice2, len(bigSlice2), cap(bigSlice2))
	bigSlice1[0] = "C#"
	fmt.Println("\nbigSlice1 = ", bigSlice1)
	fmt.Println("bigSlice2 = ", bigSlice2)
	// output:
	// bigSlice1 = [C C++ Java], len = 3, cap = 10
	// bigSlice2 = [C C++ Java Python Ruby Go], len = 6, cap = 10
	// bigSlice1 =  [C# C++ Java]
	// bigSlice2 =  [C# C++ Java Python Ruby Go]

	// 附加到具有零值的 Slice
	// 當你將值附加到 nil 的 Slice，它分配一個新的 Slice，並返回新 Slice 的參考。
	var nilSlice []string
	fmt.Printf("nilSlice = %v, len = %d, cap = %d\n", nilSlice, len(nilSlice), cap(nilSlice))
	nilSlice = append(nilSlice, "Cat", "Dog", "Lion", "Tiger")
	fmt.Printf("nilSlice = %v, len = %d, cap = %d\n", nilSlice, len(nilSlice), cap(nilSlice))
	// output:
	// nilSlice = [], len = 0, cap = 0
	// nilSlice = [Cat Dog Lion Tiger], len = 4, cap = 4

	// 將一個 Slice 附加到另一個 Slice
	// 我們可以使用 ... 運算子將一個slice直接附加到另一個slice。
	// 該運算元將slice展開為參數列表
	combineSlice1 := []string{"Jack", "John", "Peter"}
	combineSlice2 := []string{"Bill", "Mark", "Steve"}
	endSlice := append(combineSlice1, combineSlice2...)
	fmt.Println("combineSlice1 = ", combineSlice1)
	fmt.Println("combineSlice2 = ", combineSlice2)
	fmt.Println("After appending combineSlice1 & combineSlice2 = ", endSlice)
	// output:
	// combineSlice1 =  [Jack John Peter]
	// combineSlice2 =  [Bill Mark Steve]
	// After appending combineSlice1 & combineSlice2 =  [Jack John Peter Bill Mark Steve]

	// Slice 中的 Slice =================================================================
	// Slice 可以是任何型別。它們還可以包含其他 Slice
	multiSlice := [][]string{
		{"India", "China"},
		{"USA", "Canada"},
		{"Switzerland", "Germany"},
	}
	fmt.Println("multiSlice = ", multiSlice)
	fmt.Println("length = ", len(multiSlice))
	fmt.Println("capacity = ", cap(multiSlice))
	// output:
	// multiSlice =  [[India China] [USA Canada] [Switzerland Germany]]
	// length =  3
	// capacity =  3

	// 迭代 Slice =================================================================
	// 1.使用 for 迴圈迭代 Slice
	countries := []string{"India", "America", "Russia", "England"}
	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])
	}
	// output:
	// India
	// America
	// Russia
	// England

	// 2. 使用 range 形式的 for 迴圈迭代 Slice
	primeNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	for index, number := range primeNumbers {
		fmt.Printf("PrimeNumber(%d) = %d\n", index+1, number)
	}
	// output:
	// PrimeNumber(1) = 2
	// PrimeNumber(2) = 3
	// PrimeNumber(3) = 5
	// PrimeNumber(4) = 7
	// PrimeNumber(5) = 11
	// PrimeNumber(6) = 13
	// PrimeNumber(7) = 17
	// PrimeNumber(8) = 19
	// PrimeNumber(9) = 23
	// PrimeNumber(10) = 29

	// 使用空白識別符號從 range 形式的 for 迴圈中忽略 index
	numbers := []float64{3.5, 7.4, 9.2, 5.4}
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Total Sum = %.2f\n", sum)
	// output:
	// Total Sum = 25.50
}
