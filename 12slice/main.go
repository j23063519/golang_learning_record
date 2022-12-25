package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slice in golang")

	var fruitList = []string{}
	fmt.Printf("Value of fruitList is %v\n", fruitList)
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	// output:
	// Value of fruitList is []
	// Type of fruitList is []string

	// 增加元素到切片
	fruitList = append(fruitList, "Apple", "Tomato", "Banana")
	fruitList = append(fruitList, "Mango", "Peach", "Orange")
	fmt.Printf("Value of fruitList is %v\n", fruitList)
	// output:
	// Value of fruitList is [Apple Tomato Banana Mango Peach Orange]

	// 移除切片裡的元素
	// 方式一：保留從頭到第4個元素，剩下的移除，所以 Peach Orange移除了
	fruitList = append(fruitList[:4])
	fmt.Printf("Value of fruitList is %v\n", fruitList)
	// output:
	// Value of fruitList is [Apple Tomato Banana Mango]
	// 方式二:保留第二個到最後的元素，所以Apple Tomato不見了
	fruitList = append(fruitList[2:])
	fmt.Printf("Value of fruitList is %v\n", fruitList)
	// output:
	// Value of fruitList is [Banana Mango]
	// 方式三:指定保留的區間元素
	fruitList = append(fruitList[0:1])
	fmt.Printf("Value of fruitList is %v\n", fruitList)
	// output:
	// Value of fruitList is [Banana]

	// 用 make 函數建立一個初始切片，並給定初始長度
	makeSlice := make([]int, 4)
	fmt.Printf("Value of makeSlice is %v\n", makeSlice)
	fmt.Printf("Type of makeSlice is %T\n", makeSlice)
	// output:
	// Value of makeSlice is [0 0 0 0]
	// Type of makeSlice is []int

	// 初始賦值時，不可以超出長度
	// makeSlice[0] = 0
	// makeSlice[1] = 1
	// makeSlice[2] = 2
	// makeSlice[3] = 3
	// makeSlice[4] = 4
	// fmt.Printf("Value of makeSlice is %v\n", makeSlice)
	// output:
	// panic: runtime error: index out of range [4] with length 4

	// 但是可以用 append 方式增加或是減少 元素
	// 增加元素
	makeSlice[0] = 0
	makeSlice[1] = 1
	makeSlice[2] = 2
	makeSlice[3] = 3
	makeSlice = append(makeSlice, 4, 5, 6, 7)
	fmt.Printf("Value of makeSlice is %v\n", makeSlice)
	// output:
	// Value of makeSlice is [0 1 2 3 4 5 6 7]
	// 減少元素
	makeSlice = append(makeSlice[:5])
	fmt.Printf("Value of makeSlice is %v\n", makeSlice)
	// output:
	// Value of makeSlice is [0 1 2 3 4]

	// 針對切片裡的元素做排序
	makeSlice = append(makeSlice, 0, 1, 6, 7)
	sort.Ints(makeSlice)
	fmt.Printf("Value of makeSlice is %v\n", makeSlice)
	// output:
	// Value of makeSlice is [0 0 1 1 2 3 4 6 7]

	// 判斷切片裡的元素是否有排序
	fmt.Printf("Whether makeSlice is sort or not , result is %v\n", sort.IntsAreSorted(makeSlice))
	// output:
	// Whether makeSlice is sort or not , result is true

	// how to remove a value from slices based on index
	var courses = []string{"java", "python", "golang", "ruby", "javascript", "typescript"}
	fmt.Println(courses)
	// If I want to remove golang this element，I can do:
	var index = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	// output:
	// [java python golang ruby javascript typescript]
	// [java python ruby javascript typescript]

	// 進階一點：
	// 官方介紹
	// An array has a fixed size.
	// A slice, on the other hand, is a dynamically-sized,
	// flexible view into the elements of an array.
	// 簡單來說，為了讓陣列能有動態變換長度的功能，進而產生基於陣列(array)的新型別切片(slice)
	// 把切片(slice)拆開看，它包含了三個要素：
	// 長度(length):切片有多少個元素
	// 容量(capacity):能夠容納多少容量
	// 指標(pointer to array):該切片(slice)所指向的陣列(array)

	// 宣告切片
	// 一
	// 當設定切片長度為5時，雖然沒設定容量，它也會把容量大小設定成跟長度一樣的值
	var aList = make([]int, 5)
	fmt.Printf("Value of aList: %v\n", aList)
	fmt.Printf("Length of aList: %v\n", len(aList))
	fmt.Printf("Capacity of aList: %v\n", cap(aList))
	fmt.Printf("Is aList length is zero or not ? ,resault is %v\n", len(aList) == 0)
	fmt.Printf("Is aList equals to nil or not ? ,resault is %v\n", aList == nil)
	// output:
	// Value of aList: [0 0 0 0 0]
	// Length of aList: 5
	// Capacity of aList: 5
	// Is aList length is zero or not ? ,resault is false
	// Is aList equals to nil or not ? ,resault is false

	// 二
	// 同時設定切片長度及容量
	var bList = make([]int, 5, 10)
	fmt.Printf("Value of bList: %v\n", bList)
	fmt.Printf("Length of bList: %v\n", len(bList))
	fmt.Printf("Capacity of bList: %v\n", cap(bList))
	fmt.Printf("Is bList length is zero or not ? ,resault is %v\n", len(bList) == 0)
	fmt.Printf("Is bList equals to nil or not ? ,resault is %v\n", bList == nil)
	// output:
	// Value of bList: [0 0 0 0 0]
	// Length of bList: 5
	// Capacity of bList: 10
	// Is bList length is zero or not ? ,resault is false
	// Is bList equals to nil or not ? ,resault is false

	// 三
	// 初始化切片(slice)
	// 從以下結果可以得知，初始化的切片，裡面是沒有任何元素，長度為零，容量為零，但是不是零值(nil)是空切片[]
	var cList = []int{}
	fmt.Printf("Value of cList: %v\n", cList)
	fmt.Printf("Length of cList: %v\n", len(cList))
	fmt.Printf("Capacity of cList: %v\n", cap(cList))
	fmt.Printf("Is cList length is zero or not ? ,resault is %v\n", len(cList) == 0)
	fmt.Printf("Is cList equals to nil or not ? ,resault is %v\n", cList == nil)
	// output:
	// Value of cList: []
	// Length of cList: 0
	// Capacity of cList: 0
	// Is cList length is zero or not ? ,resault is true
	// Is cList equals to nil or not ? ,resault is false

	// 四
	// 尚未初始化
	// 由以下結果可以知道，尚未實體化的切片，除了沒有任何元素，長度為零，容量為零之外還是零值(nil)
	var dList []int
	fmt.Printf("Value of dList: %v\n", dList)
	fmt.Printf("Length of dList: %v\n", len(dList))
	fmt.Printf("Capacity of dList: %v\n", cap(dList))
	fmt.Printf("Is dList length is zero or not ? ,resault is %v\n", len(dList) == 0)
	fmt.Printf("Is dList equals to nil or not ? ,resault is %v\n", dList == nil)
	// output:
	// Value of dList: []
	// Length of dList: 0
	// Capacity of dList: 0
	// Is dList length is zero or not ? ,resault is true
	// Is dList equals to nil or not ? ,resault is true

	// slice 要注意的地方:
	// 狀況一
	// 若長度設定零，但是容量設定10時，會發生:
	var bigCapSlice = make([]int, 0, 10)
	fmt.Println(bigCapSlice)
	var nonZeroSlice = append(bigCapSlice, 1, 2)
	fmt.Println(nonZeroSlice)
	_ = append(bigCapSlice, 99, 88, 77)
	fmt.Println(nonZeroSlice)
	// output:
	// []
	// [1 2]
	// [99 88]

	// 狀況二
	var smallCapSlice = make([]int, 0, 2)
	fmt.Println(smallCapSlice)
	var nonZeroSlice2 = append(smallCapSlice, 1, 2)
	fmt.Println(nonZeroSlice2)
	_ = append(smallCapSlice, 99, 88, 77)
	fmt.Println(nonZeroSlice2)
	// output:
	// []
	// [1 2]
	// [1 2]

	// 根據上述兩個狀況，可以知道容量的大小會影響切片的變化，若切片容量足夠，會允許賦值，反之則不行
	// 而之所以nonZeroSlice及nonZeroSlice2兩個切片都會因_ = append(smallCapSlice, 99, 88, 77)，有所變化時，是因爲他們都是指向同個陣列(array)
	// 切片的要素之一，指標

	// detail
	detail()
}
