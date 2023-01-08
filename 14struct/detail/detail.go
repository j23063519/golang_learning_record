package detail

import (
	"fmt"
	"reflect"
)

func init() {
	// 較詳細的解說與應用
	type Persons struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	fmt.Printf("Persons: %+v\n", Persons{})
	// output:
	// Persons: {Name:}

	// 匿名 struct
	unKnown := struct {
		Name string `json:"name"`
	}{
		Name: "unKnown",
	}
	fmt.Printf("unKnown: %+v\n", unKnown)
	// output:
	// unKnown: {Name:unKnown}

	// 三種宣告 Persons struct 方式
	// 使用 new syntax：第二種和第三種寫法是一樣的
	// 1.
	var person1 *Persons
	// 2.
	person2 := &Persons{}
	// 3.
	person3 := new(Persons)
	fmt.Printf("person1: %+v\n", person1)
	fmt.Printf("person2: %+v\n", person2)
	fmt.Printf("person3: %+v\n", person3)
	// output:
	// person1: <nil>
	// person2: &{Name:}
	// person3: &{Name:}
	// structs 是在 GO 中的一種資料型態，它就類似於 javascript 裡的物件 (object)

	// 定義與使用基本的 struct
	// 1.根據資料輸入的順序，決定誰是 Name 和 Age
	mike := Persons{"mike", 13}
	fmt.Printf("mike: %+v\n", mike)
	// 2.直接取得 struct 的 pointer
	ann := &Persons{"ann", 12}
	fmt.Printf("ann: %+v\n", ann)
	// 3.key : value
	tom := Persons{
		Name: "tom",
		Age:  18,
	}
	fmt.Printf("tom: %+v\n", tom)
	// 先宣告再賦值
	var jay Persons
	jay.Name = "jay"
	jay.Age = 1
	fmt.Printf("jay: %+v\n", jay)
	// output:
	// mike: {Name:mike Age:13}
	// ann: &{Name:ann Age:12}
	// tom: {Name:tom Age:18}
	// jay: {Name:jay Age:1}

	// 當 pointer 指到的是 struct 時
	// 一般來說，若想要透過 struct pointer (&v) 來修改該 struct 中的屬性，需要先解出其值 (*p) 再 (*p).X = 10，但是這樣做太麻煩，因此 go 允許開發者直接使用 p.X 直接賦值。
	ray := &Persons{
		Name: "ray",
	}
	ray.Age = 10
	fmt.Printf("ray: %+v\n", ray)
	// output:
	// ray: &{Name:ray Age:10}

	// 另外，使用 struct pointer 時，才可以修改到原本的物件，否則會複製到一個新的
	type Car struct {
		Name string
		Age  int
	}

	// 沒使用 pointer
	toyota := Car{
		Name: "toyota",
		Age:  18,
	}
	toyotaOld := toyota
	toyotaOld.Age = 30
	fmt.Printf("toyota: %+v\n", toyota)
	fmt.Printf("toyotaOld: %+v\n", toyotaOld)
	// output:
	// toyota: {Name:toyota Age:18}
	// toyotaOld: {Name:toyota Age:30}
	// 由上述可知，當 toyotaOld 的 Age 改變時，並未連同 toyota 的 Age 一起變動

	// 使用 pointer
	mini := &Car{
		Name: "mini",
		Age:  5,
	}
	miniOld := mini
	miniOld.Age = 10
	fmt.Printf("mini: %+v\n", mini)
	fmt.Printf("miniOld: %+v\n", miniOld)
	// output:
	// mini: &{Name:mini Age:10}
	// miniOld: &{Name:mini Age:10}
	// 由上述可知，當 miniOld 的 Age 改變時，會連同 mini 的 Age 一起變動

	// nested struct
	// 在struct 裡面關聯另一個 struct
	type Info struct {
		Mail  string
		Phone int
	}
	type Man struct {
		Name string
		Age  int
		Inf  Info
	}

	ted := Man{
		Name: "ted",
		Age:  18,
		Inf: Info{
			Mail:  "ted@gmail.com",
			Phone: 1,
		},
	}
	ken := Man{
		Name: "ken",
		Age:  18,
	}
	fmt.Printf("ted: %+v\n", ted)
	fmt.Printf("ken: %+v\n", ken)
	// output:
	// ted: {Name:ted Age:18 Inf:{Mail:ted@gmail.com Phone:1}}
	// ken: {Name:ken Age:18 Inf:{Mail: Phone:0}}

	// struct field tag (meta-data)
	// struct field tag 會在 struct 後面的 key 後面，使用 backtick 來表示，例如 json:"city"
	type Weather struct {
		City string `json:"city"`
		Temp int    `json:"temp,omitempty"`
		time int    `json:"-"`
	}
	// field tag 中還可以放其他關鍵字
	// omitempty:該關鍵字是表示若該欄位沒值的話，就不會顯示該欄位名稱
	// -:表示忽略該欄位，在 Marshal 時該欄位不會出現在 JSON 中，Unmarshal 該欄位也不會被處理

	// 除了上述功能外，也可以使用 field tag 作為註解用
	type Bird struct {
		Name string `Name:"名字"`
		Age  int    `Age:"年紀"`
	}
	eagle := Bird{
		Name: "Eagle",
		Age:  18,
	}
	showEagle := reflect.TypeOf(eagle).Field(0)
	fmt.Printf("%v %v\n", showEagle.Tag, showEagle.Name)
	// output:
	// Name:"名字" Name

	// anonymous fields
	// 在 struct 中不一定要為欄位建立名稱，可以直接使用 data type，GO 會以 data type 作為欄位名稱
	type NoKeyName struct {
		int
		bool
		string
	}
	noName := NoKeyName{1, true, "Hi"}
	fmt.Printf("noName: %+v\n", noName)
	// output:
	// noName: {int:1 bool:true string:Hi}

	// function fields
	// struct 中的欄位也可以是函式
	type ShowName func(string, string) string
	type House struct {
		Name    string
		GetName ShowName
	}
	oneOone := House{
		Name: "101",
		GetName: func(Name, Where string) string {
			return "Hers is " + Where + " Hi " + Name
		},
	}

	res := oneOone.GetName("101", "taiwan")
	fmt.Println(res)
	// Hers is taiwan Hi 101

	// Promoted fields / Anonymous / Embedded
	//
}
