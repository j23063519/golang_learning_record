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

	// 若 變數是 int 型態，就得要 int指標來指向
	// a 是 int型態 的變數，b 就需要是 int的指標
	// var b *string
	// b = &a 會報錯 => cannot use &a (value of type *int) as *string value in assignment
	var a int = 10
	var b *int
	b = &a
	fmt.Printf("Type of a is: %T\n", a)
	fmt.Printf("Value of a is: %v\n", a)
	fmt.Printf("Type of b is: %T\n", b)
	fmt.Printf("Value of b is: %v\n", b)
	fmt.Printf("Value of *b is: %v\n", *b)
	// output:
	// Type of a is: int
	// Value of a is: 10
	// Type of b is: *int
	// Value of b is: 0x1400011c008
	// Value of *b is: 10

	// 透過`*向址取值`的方式來改變變數裡面的內容值
	*b = 12
	fmt.Printf("Value of a is: %v\n", a)
	fmt.Printf("Value of *b is: %v\n", *b)
	// output:
	// Value of a is: 12
	// Value of *b is: 12

	y := &b
	fmt.Printf("Value of **y is: %v\n", **y) //到y取值(到b本身的位址取值) 再取值，意即變數a

	// *y 是 &b 亦即是 取得址
	// **y 就是 a 就是 取值
	**y = 15
	fmt.Printf("Value of a is: %v\n", a)
	fmt.Printf("Value of *b is: %v\n", *b)
	fmt.Printf("Value of **y is: %v\n", **y)
	// output:
	// Value of a is: 15
	// Value of *b is: 15
	// Value of **y is: 15

	// swap function 交換數值 =================================================================================================
	c, d := 10, 20
	fakeSwap(c, d)
	fmt.Printf("Value of c is: %v\n", c)
	fmt.Printf("Value of d is: %v\n", d)
	// 雖然這裡沒變，但其實在函式裡面是有變化的
	// 但是交換不成的原因是因為交換函式所宣告的值
	// output:
	// Value of c is: 10
	// Value of d is: 20

	// 這樣就能確實交換數值
	e, f := &c, &d
	trueSwap(e, f)
	fmt.Printf("Value of e is: %v\n", *e)
	fmt.Printf("Value of f is: %v\n", *f)
	// output:
	// Value of e is: 20
	// Value of f is: 10

	// struct 物件上的指標 =================================================================================================
	// 這邊想要針對 TestStruct這個物件裡面的屬性 Name做變更，以下分別是無效及有效的範例
	// 無效的範例
	var fake1 = TestStruct{Name: "Hi"}
	fake1.showName()
	var fake2 = TestStruct{Name: "Hi 1"}
	fake2.fakeChangeName("Hi 2")
	fake2.showName()
	// output:
	// 測試 Hi 名稱
	// 測試 Hi 1 名稱 (Hi 2 沒顯示)

	// 有效的範例
	fake2.trueChangeName("Hi 2")
	fake2.showName()
	// output:
	// 測試 Hi 2 名稱

	// struct 物件上的指標 new object function ============================================================================
	// 若要自製一個可以產生新物件的函式，可以透過 struct 物件 pointer 的概念去實作，如同New這個方法一樣
	n := &newStruct{"物件"}
	fmt.Printf("Value of newStruct is: %v\n", n)
	fmt.Printf("Pointer address of newStruct is: %v\n", &n)
	// output:
	// Value of newStruct is: &{物件}
	// Pointer address of newStruct is: 0x1400009e020

	// 客製函式
	n1 := newStructMethod("")
	fmt.Printf("Value of newStruct is: %v\n", n1)
	fmt.Printf("Pointer address of newStruct is: %v\n", &n1)
	// output:
	// Value of newStruct is: &{}
	// Pointer address of newStruct is: 0x14000120028

	// 客製函式
	n2 := newStructMethod("新物件")
	fmt.Printf("Value of newStruct is: %v\n", n2)
	fmt.Printf("Pointer address of newStruct is: %v\n", &n2)
	// output:
	// alue of newStruct is: &{新物件}
	// Pointer address of newStruct is: 0x14000120030

	// go 提供的 New方法
	n3 := new(newStruct)
	fmt.Printf("Value of newStruct is: %v\n", n3)
	fmt.Printf("Pointer address of newStruct is: %v\n", &n3)
	// output:
	// Value of newStruct is: &{}
	// Pointer address of newStruct is: 0x1400000e048

	// 較詳細的解釋應用
	// ptr 這個function 是在ptr.go 的 function
	ptr()
}

// 無效的範例
func fakeSwap(c int, d int) {
	newC := c
	c = d
	d = newC
	fmt.Printf("FakeSwap value of c is: %v\n", c)
	fmt.Printf("FakeSwap value of d is: %v\n", d)
	// output:
	// FakeSwap value of c is: 20
	// FakeSwap value of d is: 10
}

// 有效的範例
func trueSwap(c *int, d *int) {
	newC := *c
	*c = *d
	*d = newC
}

// TestStruct 物件
type TestStruct struct {
	Name string
}

// TestStruct 物件 裡的方法 showName
func (t TestStruct) showName() {
	fmt.Println("測試", t.Name, "名稱")
}

// TestStruct 物件 裡的方法 fakeChangeName (無效)
func (t TestStruct) fakeChangeName(newName string) {
	t.Name = newName
}

// TestStruct 物件 裡的方法 trueChangeName (有效)
func (t *TestStruct) trueChangeName(newName string) {
	t.Name = newName
}

// NewStruct 物件
type newStruct struct {
	name string
}

// newStruct 物件 裡的方法 newStructMethod
func newStructMethod(n string) *newStruct {
	// tip:
	// 為何 不是這樣寫 return &newStruct{name: n}
	// 因爲go 可以自動把 name屬性賦值n
	return &newStruct{n}
}
