package main

import (
	"fmt"
	"math"
	"reflect"
)

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width, height float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perimeter() float64 {
	return c.radius * 2
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func getArea(s shape) float64 {
	return s.area()
}

// Generic function to calculate the total area of multiple shapes of different types
func getTotalArea(shapes ...shape) float64 {
	totalArea := 0.0
	for _, v := range shapes {
		totalArea += v.area()
	}
	return totalArea
}

// Interface types can also be used as fields
type MyDrawing struct {
	shapes  []shape
	bgColor string
	fgColor string
}

func (md MyDrawing) Area() float64 {
	total := 0.0
	for _, v := range md.shapes {
		total += v.area()
	}
	return total
}

func main() {
	// interface
	c1 := circle{4.5}
	r1 := rect{10, 10}
	shapes := []shape{c1, r1}

	for _, shape := range shapes {
		fmt.Printf("area: %v\n", shape.area())
		// output:
		// area: 63.61725123519331
		// area: 100
		fmt.Printf("area: %v\n", getArea(shape))
		// output:
		// area: 63.61725123519331
		// area: 100
	}

	// 使用介面型別作為函數的參數
	fmt.Println("Total area = ", getTotalArea(circle{4.5}, rect{10, 10}, rect{10, 5}))
	// output:
	// Total area =  213.6172512351933

	// 將介面型別用於欄位
	drawing := MyDrawing{
		shapes: []shape{
			circle{4.5}, rect{10, 10}, rect{10, 5},
		},
		bgColor: "red",
		fgColor: "white",
	}
	fmt.Println("Drawing", drawing)
	fmt.Println("Drawing Area = ", drawing.Area())
	// output:
	// Drawing {[{4.5} {10 10} {10 5}] red white}
	// Drawing Area =  213.6172512351933

	// interface 介面
	// 介面的意思是開放了功能出來，這些功能與內部程式碼實作的部分完全無關，你不知道也可以用
	// 像是你不需要知道車子的所有製造細項，你只要如何使用即可
	// 在 Go 語言中，interface 就是任意值
	var anyTypeValue interface{}
	fmt.Printf("Value: %v, Type: %v\n", anyTypeValue, reflect.TypeOf(anyTypeValue))
	anyTypeValue = 123
	fmt.Printf("Value: %v, Type: %v\n", anyTypeValue, reflect.TypeOf(anyTypeValue))
	anyTypeValue = false
	fmt.Printf("Value: %v, Type: %v\n", anyTypeValue, reflect.TypeOf(anyTypeValue))
	anyTypeValue = "Hi"
	fmt.Printf("Value: %v, Type: %v\n", anyTypeValue, reflect.TypeOf(anyTypeValue))
	// output:
	// Value: <nil>, Type: <nil> 因為還沒實體化，所以為 nil
	// Value: 123, Type: int
	// Value: false, Type: bool
	// Value: Hi, Type: string

	kitty := Kitten{
		Name: "kitty",
	}
	kitty.eat()
	// output:
	// kitty eating

	// 斷言：表示宣稱的意思
	// 型別斷言表示存取介面實體化之後值的方法，但是如果給錯型別會引發
	var i interface{} = "Hi"
	hiStr := i.(string)
	// hiInt := i.(int)
	fmt.Println(hiStr)
	// fmt.Println(hiInt)
	// output:
	// Hi
	// panic: interface conversion: interface {} is string, not int

	// 若是不知道型別，可以用 ok 來判斷，轉換的型別是否成功
	if val, ok := i.(int); !ok {
		fmt.Println(val, ok)
		// output:
		// 0 false
	}

	// 除了用if判斷外，也可以用 switch 加上 type 判斷
	// i.(type) 這是比較特殊的用法，比較常見的是與 switch 搭配使用
	switch v := i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("unknown", v)
	}
	// output:
	// string

	// 稍微細節介紹
	detail()
}

// interface 除了任意值，還有一個更廣泛的用法，可以定義 func 在 interface 裡面，
// 用這樣的方式可以讓程式更抽象化，更有彈性
// 可以 定義一個 Anamal 動物 的 interface 裡面有 eat func
type Anamal interface {
	eat()
}

type Kitten struct {
	Name string
}

func (k Kitten) eat() {
	fmt.Println(k.Name, "eating")
}
