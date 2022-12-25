package main

import "fmt"

func detail() {
	// map 是沒有順序的鍵值對集合，鍵在 map 是唯一的，但值不一定。
	// map 資料結構用於基於鍵的快速尋找，取得和刪除資料。它是最常用的資料結構之一。

	// 宣告 map
	// var m map[KeyType]ValueType
	// 宣告一個 string 的鍵及 int 的值 的map
	// var m map[string]int

	// map的零值是nil，且nil的map並沒有鍵
	// 此外，任何嘗試將鍵加到 nil 的 map 的行為將導致執行時期錯誤。
	var a map[string]int
	fmt.Printf("value: %+v\n", a)
	if a == nil {
		fmt.Println("a is nil")
	}
	// output:
	// value: map[]
	// a is nil

	// panic: assignment to entry in nil map
	// a["one hundred"] = 100
	// 因此，必需在加入項目之前初始化 map。

	// 初始化 map =================================
	// 1.使用 make()函數初始化 map
	var b = make(map[string]int)
	fmt.Printf("value: %+v\n", b)

	if b == nil {
		fmt.Println("b is nil")
	} else {
		fmt.Println("b is not nil")
	}

	b["keke"] = 99
	fmt.Printf("value: %+v\n", b)
	// output:
	// value: map[]
	// b is not nil
	// value: map[keke:99]

	// 2.使用 map 定數初始化 map
	// map 定數是使用某些資料初始化 map的一種方法
	var c = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5, // Comma is necessary
	}
	fmt.Printf("value: %+v\n", c)
	// output:
	// value: map[five:5 four:4 one:1 three:3 two:2]

	// 加入項目(鍵值對)到map =================================
	var d = make(map[string]string)

	// Adding keys to a map
	d["one"] = "1"
	d["two"] = "2"
	d["three"] = "3"
	d["four"] = "4"
	d["five"] = "5"
	fmt.Printf("value: %+v\n", d)

	d["five"] = "ha"
	fmt.Printf("value: %+v\n", d)
	// output:
	// value: map[five:5 four:4 one:1 three:3 two:2]
	// value: map[five:ha four:4 one:1 three:3 two:2]

	// 取得 map 中特定鍵所關聯的值 =================================
	// 你可以使用語法 m[key] 來取得分派給 map 中鍵的值
	var e = map[string]string{
		"John":  "+33-8273658526",
		"Steve": "+1-8579822345",
		"David": "+44-9462834443",
	}

	var newE = e["Steve"]
	fmt.Println("Steve's Mobile No : ", newE)

	// If a key doesn't exist in the map, we get the zero value of the value type
	newE = e["Jack"]
	fmt.Println("Jack's Mobile No : ", newE)
	// output:
	// Steve's Mobile No :  +1-8579822345
	// Jack's Mobile No :
	// 在上面的範例中，由於鍵 "Jack" 不存在於 map 中，我們得到 map 的值型別的零值。由於 map 的值型別是 string，因此我們得到 ""。

	// map 是參考型別 =================================
	// map 是參考型別，當你將 map 分派給新變數時，他們都參考相同的底層資料結構。因此，一個變數完成的更改另一個變數也會如此。
	var sameMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	var sameMap2 = sameMap
	fmt.Println("sameMap = ", sameMap)
	fmt.Println("sameMap2 = ", sameMap2)

	sameMap2["ten"] = 10
	fmt.Println("\nsameMap = ", sameMap)
	fmt.Println("sameMap2 = ", sameMap2)
	// output:
	// sameMap =  map[five:5 four:4 one:1 three:3 two:2]
	// sameMap2 =  map[five:5 four:4 one:1 three:3 two:2]

	// sameMap =  map[five:5 four:4 one:1 ten:10 three:3 two:2]
	// sameMap2 =  map[five:5 four:4 one:1 ten:10 three:3 two:2]
}
