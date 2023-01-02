package main

import (
	"fmt"
	"strIntro/model"
)

type User struct {
	// If attributes have same data type , we can put them together
	Name, Email string
	Status      bool
	Age         int
}

func main() {
	fmt.Printf("Struct introduction\n")

	fmt.Printf("man's details are %+v\n", User{})
	// output:
	// man's details are {Name: Email: Status:false Age:0}

	man := User{
		Name: "Man",
		Age:  18,
	}
	fmt.Printf("man's details are %+v\n", man)
	// output:
	// man's details are {Name:Man Email: Status:false Age:18}

	// 說明：
	// struct 是使用者定義的型別，包含已命名欄位/屬性的集合。它用於將相關資料分組在一起形成一個單位。

	// 定義 struct 型別
	// type Person struct {
	// 	FirstName string
	// 	LastName  string
	// 	Age       int
	// }
	// type 關鍵字引進了新型別。其後是型別(Person)的名稱和關鍵字 struct，表示我們正在定義一個 struct。struct 包含了一個大括號內的 欄位 列表。
	// 若有相同型別的屬性可以放入同欄位
	// type Person struct {
	// 	FirstName, LastName string
	// 	Age                 int
	// }

	// 宣告 struct 型別的變數
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	var p Person
	fmt.Printf("Person %+v\n", p)
	// output:
	// Person {FirstName: LastName: Age:0}
	// 上述程式碼有一個類型為Person的變數p，預設情況下為零值，其欄位 FirstName 和 LastName 設定為""，且 Age 設定為0

	// 初始化 struct
	// 你可以使用 struct 定數來初始化一個 struct 型別的變數:
	var human = Person{"John", "Smith", 24}
	fmt.Printf("human %+v\n", human)
	// output:
	// human {FirstName:John LastName:Smith Age:24}

	// 另外，要記得不能初始化部分欄位
	// var wrongPerson = Person{"Naem"}
	// fmt.Printf("wrongPerson %+v\n", wrongPerson)
	// too few values in Person{…}

	// 可以使用 name:value 語法 初始化部分欄位
	// 使用該語法時，欄位順序無關緊要
	var woman = Person{FirstName: "Lisa", Age: 16}
	fmt.Printf("woman %+v\n", woman)
	// output:
	// woman {FirstName:Lisa LastName: Age:16}

	// 可以用換行符號分隔多個欄位，提高可讀性(需用逗號結尾)
	var other = Person{
		LastName:  "John",
		FirstName: "Ku",
		Age:       18,
	}
	fmt.Printf("other %+v\n", other)
	// output:
	// other {FirstName:Ku LastName:John Age:18}

	// 存取 struct 的欄位
	// 可以使用點(.)運算元存取 struct 的各個欄位
	type Car struct {
		Name, Modelm, Color string
		Weight              int
	}
	var mini = Car{
		Name:   "Ferrari",
		Modelm: "GTC4",
		Color:  "red",
		Weight: 100,
	}

	fmt.Println("Car Name: ", mini.Name)
	fmt.Println("Car Color: ", mini.Color)
	// output:
	// Car Name:  Ferrari
	// Car Color:  red

	// 重新賦值
	mini.Color = "Black"
	fmt.Println("Car: ", mini)
	// output:
	// Car:  {Ferrari GTC4 Black 100}

	// 指向 struct 的指標
	// 可以使用 &運算元取得指向 struct 的指標
	type Sturdent struct {
		RollNumber int
		Name       string
	}
	s := Sturdent{11, "Jack"}
	ps := &s
	fmt.Println(ps)
	// output:
	// &{11 Jack}

	// 可以使用下述兩個方式：取得 Name 值
	fmt.Println((*ps).Name)
	fmt.Println(ps.Name)
	// Jack
	// Jack

	ps.RollNumber = 31
	fmt.Println(ps)
	// output:
	// &{31 Jack}

	// 使用內建的 new()函數建立一個 struct 並取得指向他的指標
	// 也可以使用內建的 new() 函數來建立 struct 的實體。new() 函數會分配足夠的記憶體來符合所有 struct 的欄位，將他們各自設為零值，並返回指向新分配的 struct 的指標
	type Employee struct {
		Id   int
		Name string
	}
	pEmp := new(Employee)
	pEmp.Id = 1
	pEmp.Name = "Noron"
	fmt.Println(pEmp)
	// output:
	// &{1 Noron}

	// 匯出與未匯出的 struct 和 struct 欄位
	// 任何以大寫字母開頭的 struct 型別都會被匯出，並且可以從外部的 package 中存取。同樣的，任何以大寫字母開頭的 struct 欄位都會被匯出。
	// 所以以小寫字母開頭的名稱只能在同一package 中使用。
	// $GOPATH/src
	// 		example
	//   		main
	//     			main.go
	//   		model
	//     			address.go
	//     			customer.go
	c := model.Customer{
		Id:   1,
		Name: "Jay Chuo",
	}
	// c.married = true
	// c.married undefined (type model.Customer has no field or method married)

	// a := model.address{}
	// address not exported by package model
	fmt.Println("Programmer = ", c)
	// output:
	// Programmer =  {1 Jay Chuo {     0} false}
	// 由上述結果可知，address 和 married 是未匯出的，無法從 main package 中存取。

	// struct 是值型別
	// struct 是值型別，當將一個 struct 變數 分派給另一個變數時，會建立並分派一個新的 struct 副本。同樣的，當將 struct 傳遞給另一個函數時，該函數將取得自己的 struct 副本。
	type Point struct {
		X float64
		Y float64
	}
	p1 := Point{X: 10, Y: 20}
	// A copy of the struct `p1` is assigned to `p2`
	p2 := p1
	fmt.Println("p1 = ", p1)
	fmt.Println("p2 = ", p2)
	// output:
	// p1 =  {10 20}
	// p2 =  {10 20}

	p2.X = 15
	fmt.Println("\nAfter modifying p2:")
	fmt.Println("p1 = ", p1)
	fmt.Println("p2 = ", p2)
	// output:
	// After modifying p2:
	// p1 =  {10 20}
	// p2 =  {15 20}
	// 由上述結果可知：p2的屬性X重新賦值後，並不會連帶改變 p1的屬性X

	// struct 等式
	// 如果兩個 struct 變數的所有對應欄位都相等，則他們相等
	p3 := Point{1.1, 2.2}
	p4 := Point{1.1, 2.2}
	if p3 == p4 {
		fmt.Println("p3 等於 p4")
	} else {
		fmt.Println("p3 不等於 p4")

	}
	// output:
	// p3 等於 p4

	p5 := Point{1.1, 2.2}
	p6 := Point{1.2, 2.2}
	if p5 == p6 {
		fmt.Println("p5 等於 p6")
	} else {
		fmt.Println("p5 不等於 p6")

	}
	// output:
	// p5 不等於 p6
}
