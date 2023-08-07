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

	// Interface 會被隱性的實作
	var itf Itf = TpeImpItf{"implement"}
	itf.M()
	// output:
	// S: implement

	// 同時符合多個 interface 的 type
	// 。MutiItf type 同時實作了，Hand 和 Foot interface
	// 。可以把 MutiItf 指派給 Hand 和 Foot interface 所建立的變數
	// 。Hand interface 所建立的變數 hit 可以使用 hit.Grab() 的方法;Foot interface 所建立的變數 fit 可以使用 fit.Kick() 的方法
	// 。雖然變數 hit 和 fit 的 dynamic type 都是 MutiItf，但底層的 Static Tpye 不同，前者是 Hand interface、後者是 Foot interface，因此雖然他們 dynamic type 都是 MutiItf type，但是並不能使用 hit.Kick() 和 fit.Grab()
	mitf := MutiItf{baseScore: 5}
	var hit Hand = mitf
	var fit Foot = mitf
	fmt.Printf("\nBase Grap Score of hit: %v\n", hit.Grab())
	fmt.Printf("Hand (%T, %v)\n", hit, hit)
	fmt.Printf("Base Grap Score of fit: %v\n", fit.Kick())
	fmt.Printf("Foot (%T, %v)\n", fit, fit)
	// output:
	// Base Grap Score of hit: 100
	// Hand (main.MutiItf, {5})
	// Base Grap Score of fit: 150
	// Foot (main.MutiItf, {5})

	// 型別斷言 Type Assertions
	// 雖然 hit 和 fit 的 dynamic type 都是 MutiItf，但因為他們的底層 static type 並不一樣，hit 的是 Hand、fit 的是 Foot，因此不能呼叫 hit.Kick() 和 fit.Grab()的方法。為了避免呼叫到底底層並不對應的 static type 可以使用 type assertion 的方式:
	// val := i.(Type) // 得到實作 interface 的 type 的值
	// i: type interface
	// Type: 實作該 interface 的 type
	var anotherHit Hand = MutiItf{baseScore: 40}
	// 原本不能執行 anotherHit.Kick()，但把 anotherHit 轉換成 MutiItf 後，得到 newHit 即可使用 newHit.Kick()
	newHit := anotherHit.(MutiItf)
	fmt.Println("\nGrab score:", newHit.Grab())
	fmt.Printf("Hand (%T, %v)\n", newHit, newHit)
	fmt.Println("Kick score:", newHit.Kick())
	fmt.Printf("Foot (%T, %v)\n", newHit, newHit)
	// output:
	// Grab score: 800
	// Hand (main.MutiItf, {40})
	// Kick score: 1200
	// Foot (main.MutiItf, {40})

	// 若是 MutiItf 並沒有實作 Foot 或是 Hand interface 方法 或是 MutiItf 有實作 Foot 或是 Hand interface 方法但是並沒有實際賦值都會報錯，為了避免這種狀況可以使用：
	// val, ok := i.(Type)
	// ok: 當 i 有 dynamic type 和 dynamic value 時，ok 會是 true，
	// val 則會是 dynamic value;否則 ok 會是 false，value 會是該 type 的 zero value

	// Type Assertions 除了用來確保一個 interface 是否有 dynamic value / concrete value 之外，也可以用來將一個變數從原本的interface 轉變成另一個 interface
	var JohnPers Pers = Empl{"John", "Albert", 2000}
	fmt.Printf("full name: %v \n", JohnPers)
	JohnSalar := JohnPers.(Salar)
	fmt.Printf("salar: %v \n", JohnSalar.getSalary())
	// output:
	// full name: {John Albert 2000}
	// salar: 2000

	// Embedding interfaces
	// embedded，嵌入:俗稱golang 的繼承，但不是真的繼承(背後的機制和程式語言處理方式不同)。而行為可以做到類似繼承的特點。
	// 在 go 中，我們可以將兩個 interface 組合成一個新的 interface
	type ShapPart interface {
		Area() float64
	}
	type ObjectPart interface {
		Volume() float64
	}
	type MaterialCombine interface {
		ShapPart
		ObjectPart
	}
	// MaterialCombine interface 是由 ShapPart 和 ObjectPart 組合而成

	// interface 試圖解決的問題
	// 問題：共同相同邏輯但帶入不同型別參數得函式
	// 如果某一個函式內部運作邏輯相同，我還需要因參數型別不同而撰寫不同的 function呢？
	// interface 的使用
	// 在程式中的任何 type，只要這個 type 的函式有符合到該 interface 的定義，就可以歸類到該 interface 底下。
	// 範例一：polymorphism
	// 在這個例子中，在這個例子中，square type 的方法 area() 因為符合 shape interface 的定義，所以 square type 也一併被歸類在 shape interface 中（polymorphism）：
	// // STEP 1：定義 Shape 這個 interface
	// type Shape interface {
	//     area() int
	// }

	// // STEP 2：定義 square type
	// type Square struct {
	//     sideLength int
	// }

	// // STEP 3：定義 area 這個 Square type 的 methods
	// // 因為 Square type 的 area method 符合 Shape interface 的規範
	// // 所以 Square type 同樣屬於 Shape interface
	// func (s Square) area() int {
	//     return s.sideLength * s.sideLength
	// }

	// func main() {
	//     // STEP 5：定義 Square type 的變數，ten 了符合 Square Type 之外，也同時符合 Shape Interface
	//     ten := Square{sideLength: 10}

	//     // STEP 6：printArea 中可以帶入 Shape type
	//     // 因為 tne 現在同時屬於 Shape interface，所以可以放入 printArea 這個 function
	//     printArea(ten)
	// }

	// // STEP 4：printArea 這個 function 可以帶入 Shape interface 作為參數
	//
	//	func printArea(s Shape) {
	//	    fmt.Println(s.area())
	//	}
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

type Itf interface {
	M()
}

type TpeImpItf struct {
	S string
}

// type tpe 實作了 interface Itf 的 M方法
// 但是不需宣告
func (tpe TpeImpItf) M() {
	fmt.Printf("\nS: %v\n", tpe.S)
}

type Hand interface {
	Grab() uint64
}

type Foot interface {
	Kick() uint64
}

type MutiItf struct {
	baseScore uint64
}

func (mtf MutiItf) Grab() uint64 {
	return 20 * mtf.baseScore
}

func (mtf MutiItf) Kick() uint64 {
	return 30 * mtf.baseScore
}

type Pers interface {
	getFullName() string
}

type Salar interface {
	getSalary() int
}

type Empl struct {
	firstName string
	lastName  string
	salary    int
}

func (e Empl) getFullName() string {
	return e.firstName + " " + e.lastName
}

func (e Empl) getSalary() int {
	return e.salary
}
