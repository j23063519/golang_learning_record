package main

import "fmt"

func main() {
	fmt.Println("Array in golang")

	// array needs length and type => [5]int
	var fruiltList [5]string
	fruiltList[0] = "Apple"
	fruiltList[1] = "Tomato"
	fruiltList[3] = "Peach"
	fmt.Println("Fruilt list is", fruiltList)
	fmt.Println("Fruilt list length is", len(fruiltList))
	// output:
	// Fruilt list is [Apple Tomato  Peach ]
	// Fruilt list length is 5
	// 由此可知，陣列需要先給定長度，且當有未賦值的位置會以空格顯示
	// note:
	// 若為給定長度則該型別為slice(切片)

	// 陣列定好長度後不可以超出
	// fruiltList[5] = "Peach" => invalid argument: index 5 out of bounds [0:5]

	// 除了上述方式，可以在定義型別時就賦值:
	var vegList = [10]string{"potato", "beans", "", "mushrooms"}
	fmt.Println("Vegetable list is", vegList)
	fmt.Println("Vegetable list length is", len(vegList))
	// output:
	// Vegetable list is [potato beans  mushrooms      ]
	// Fruilt list length is 10

	// 也可以不定義長度，讓go 自行判斷，使用 ...
	var autoList = [...]int{1, 2, 3, 4, 5, 6, 0, 8, 9, 10}
	fmt.Println("Auto list is", autoList)
	fmt.Println("Auto list length is", len(autoList))
	// output:
	// Auto list is [1 2 3 4 5 6 0 8 9 10]
	// Auto list length is 10

	// Array:
	// 陣列的長度是固定的，屬於原生型別(primitive type)，較少在程式使用，若想要能增長或減少長度就需要使用切片(slice)
}
