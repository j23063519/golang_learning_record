package main

import "fmt"

func main() {
	// Map即是一種利用Key-Value方式來對應的資料格式
	fmt.Printf("Map in golang\n")
	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["GO"] = "golang"
	languages["PY"] = "python"

	fmt.Println("List of all languages", languages)
	fmt.Println("GO shorts for: ", languages["GO"])
	// output:
	// List of all languages map[GO:golang JS:javascript PY:python]
	// GO shorts for:  golang

	delete(languages, "PY")
	fmt.Println("List of all languages", languages)
	// output:
	// List of all languages map[GO:golang JS:javascript]

	// loops
	for k, v := range languages {
		fmt.Printf("For key %v, value is %v\n", k, v)
	}
	// output:
	// For key JS, value is javascript
	// For key GO, value is golang

	// only show index
	for k := range languages {
		fmt.Printf("For key %v\n", k)
	}
	// output:
	// For key JS
	// For key GO

	// ignore key
	for _, v := range languages {
		fmt.Printf("Value is %v\n", v)
	}
	// output:
	// Value is javascript
	// Value is golang

	// 補充:
	// 在 Go 中，Map 也是 Key-Value pair 的組合，但是 Map 所有 Key 的資料型別都要一樣；所有 Value 的資料型別也要一樣。另外，在設值和取值需要使用 [] 而非 .。map 的 zero value 是 nil。

	// 在 map 中因為 key 本身是有資料型別的，因此只能使用 [] 來設值和取值，不能使用 .
	// 使用 value, isExist := m[key] 可以用來取值，同時判斷該 key 是否存在
	value, isExist := languages["GO"]
	fmt.Printf("Value is %v, isExist is %v\n", value, isExist)
	// output:
	// Value is golang, isExist is true
	value1, isExist1 := languages["TEST"]
	fmt.Printf("Value1 is %v, isExist1 is %v\n", value1, isExist1)
	// output:
	// Value1 is , isExist1 is false

	// Map 和 Structure 的選擇
	// Map                                                             Structure
	// 所有的 key 都需要是相同型別
	// 所有的 value 都需要是相同型別                                       value 可以是不同型別
	// Key 有被 indexed，因此可以進行疊代，列出所有的 key-value pairs        Key 沒有被 indexed，因此沒辦法透過疊代的方式列出所有的 key-value pair
	// Reference Type                                                  Value Type
	// 通常用在資料集合關聯性非常強的資料，例如 colors                        通常用在一個帶有多種屬性的東西，例如 person
	// 不用一開始就知道所有的 key 有哪些，因為可以後續新增和刪除                需要一開始就清楚定義所有的欄位

	// 其他範例
	detail()
}
