package main

import (
	"fmt"
	"math"
)

func detail() {
	// Interface 是什麼
	// Interface 有點像是藍圖，先定義好某個方法的名稱 (function name)、會接收的參數及型別(list of argument type)、會回傳的值與型別(list of return types)
	// 定義好藍圖之後，並不去管實作的細節，實作的細節會由每個型別自行定義實作(implementation)
	// 透過 interface 可以定義一系列的 method signatures 來讓 Type 透過 methods 加以實作，也是說 interface 可以用來定義 type 有哪些行為
	// 舉例來說，定義 貓 的 interface 包含方法 走路及貓叫 只要有一個 Type 他能實作 走路及貓叫，那這個 Type 就自動實作 狗這個 interface，不需額外使用 implement 關鍵字，另外，任何資料型別，只要實作該 interface 之後，都可以被視為該 interface 的 type (polymorphism)

	// Interface Tips:
	// interface 沒有被賦值之前，其 type 和 value 都會是 nil
	// interface 被賦值後，他的型別會變成實作它的 Type 的型別和值
	// interface 的變數『動態值(dynamic value)』會是實作此 interface 的 Type 的 value
	// interface 的型別『動態型別(dynamic value)』會是實作此 interface 的 Type 的型別
	// interface 沒有『靜態值(static value)』
	// interface 的『靜態型別(static type)』，則是該 interface 的本身，例如 type Shape interface{}，這個 interface 所建立的變數，其靜態型別即是 Shape
	// interface 的 dynamic type 又稱作是 concrete type，因為當我們想要存取該 interface 的型別時，他回傳的會是 dynamic value，原本的 static type 會被隱藏

	// 從下述例子可以知道，目前的 interface 因爲尚未被賦值，所以回傳的是 zero value。
	var s Shaping
	fmt.Printf("\nValue of s is %v, Type of s is %T\n", s, s)
	// output:
	// Value of s is <nil>, Type of s is <nil>

	// 如果 interface 有賦值的話，則可以看到顯示的 dynamic type 和 dynamic value 會是實作該 interface 的 type 的 method 和 value
	// 。Rect type 實作了 Shape interface
	// 。當一個 type 實作(implement)某個interface後，該 Type 產生的變數除了會是原本的 Type 外，也同時屬於該 interface type，及 polymorphism
	// 。當把 Rect 作為 Shaping interface 的值後，Shape 的 Type (dynamic type)會變成 Rect、Value(dynamic value)會變成 Rect 的值({3,5})
	// 範例一
	var s1 Shaping = Rect{3, 5}
	fmt.Printf("\nValue of s1 is %v, Type of s1 is %T\n", s1, s1)
	fmt.Println(s1.Area())
	// output:
	// Value of s1 is {3 5}, Type of s1 is main.Rect
	// 15

	// 範例二
	var i I
	i = &T{"Hello"}                // 把 type T 的值賦予給變數 i
	fmt.Printf("(%v, %T)\n", i, i) // i 的 dynamic value 是 &{Hello}、dynamic type 是 *main.T
	i.M()                          // 意思是將 type T 對應的 value (&{Hello}) 來執行 type T 對應的 M 方法
	// output:
	// (&{Hello}, *main.T)
	// T.S is Hello

	i = F(math.Pi)                 // 把 type F 的值賦予給變數 i
	fmt.Printf("(%v, %T)\n", i, i) // i 的 dynamic value 是 3.141、dynamic type 是 main.F
	i.M()                          // 意思是將 type F 對應的 value  3.141 來執行 type F 對應的 M 方法
	// output:
	// (3.141592653589793, main.F)
	// 3.141592653589793

	// Interface 的 polymorphism
	// 當一個 type 實作了 interface 後，這個 type 所建立的變數除了屬於原本的 type 之外，也屬於這個 interface 的 type，一個變數符合多個型別就稱作 polymorphism (多型)
	// 舉例來說，這裡先定義了 Salaried 這個 interface type，接著 Salary 這個 struct type 實作了 Salaried interface type，這樣行為稱作 polymorphism
	mike := Employee{
		firstName: "Mike",
		lastName:  "John",
		salary: Salary{
			1100, 40, 20,
		},
	}
	fmt.Println("Mike's salary is", mike.salary.getSalary())
	// output:
	// Mike's salary is 1160
}

// interface 沒有被賦值前，其 type 和 value 都會是 `nil`
type Shaping interface {
	Area() float64
}

// Rect 實作了 Shape interface
// Rect 所建立的變數同時會符合 Rect (Struct Type)和 Shape (Interface Type)
type Rect struct {
	w float64
	h float64
}

func (r Rect) Area() float64 {
	return r.w * r.h
}

// 對應上述範例二
type I interface {
	M()
}

// Type T 實作了 I interface
type T struct {
	S string
}

func (t T) M() {
	fmt.Println("T.S is", t.S)
}

// Type F 實作了 I interface
type F float64

func (f F) M() {
	fmt.Println(f)
}

// Interface 的 polymorphism
type Salaried interface {
	getSalary() int
}

// Salary 實作了 getSalary() 的方法，因此可以算是 Salaried type (polymorphism 多型)
type Salary struct {
	basic, insurance, allowance int
}

func (s Salary) getSalary() int {
	return s.basic + s.insurance + s.allowance
}

type Employee struct {
	firstName, lastName string
	salary              Salaried
}
