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
	// 一：在 golang 中 struct 的 field name 可以省略，沒有 field name 的 name，稱作 anonymous or embedded.
	// 這種狀況下會直接使用「Type」名稱當作 field name
	// 二：如果 nested anonymous struct 裡的欄位跟 parent struct 欄位有一樣的話，則該欄位不會被 promoted
	// 三：對於 anonymous(embedded) fields 的欄位(field)或是方法(method)都稱作promoted，他們就像一般的欄位一樣，但是不能跳過 Type 的名稱直接用 struct literals 的方式來賦值

	// 錯誤用法：不能在未明確定義 promoted fields 名稱的情況下，使用 struct literals 設值
	// record := Record{
	// 	name:     "record",
	// 	age:      29,
	// 	position: "software engineer",
	// }
	// fmt.Printf("%+v", record)
	// undeclared name: Record compiler

	// 正確：明確定義要設值的 promoted fields 名稱為何
	type Person struct {
		name string
		age  int32
	}
	type Employee struct {
		position string
	}
	type Record struct {
		Person
		Employee
	}
	record := Record{
		Person: Person{
			name: "record",
			age:  29,
		},
		Employee: Employee{
			position: "software engineer",
		},
	}

	fmt.Printf("%+v\n", record)
	// output:
	// {Person:{name:record age:29} Employee:{position:software engineer}}Struct introduction

	// 在 Promoted fields 中取值
	// 不論有沒有使用明確的 promoted fields 名稱，都可以取值
	// 根據上述範例：
	fmt.Println("name:", record.name)                           // 29
	fmt.Println("Person.name:", record.Person.name)             // 29
	fmt.Println("position:", record.name)                       // software engineer
	fmt.Println("Employee.position:", record.Employee.position) // software engineer
	// output:
	// name: record
	// Person.name: record
	// position: record
	// Employee.position: software engineer

	// 組裝式繼承
	goku := &Saiyan{
		Human: &Human{"Goku"},
		Power: 1000,
	}

	fmt.Println("name", goku.Name)
	fmt.Println("name", goku.Human.Name)
	goku.Introduce()
	goku.Human.Introduce()
	// output:
	// name Goku
	// name Goku
	// Hi, I'm Goku
	// Hi, I'm Goku

	// Interface Fields (Nested Interface)
	tim := Worker{
		firstName: "Jay",
		lastName:  "Geller",
		salary: Salary{
			1100, 50, 50,
		},
	}
	// 因為 Salary struct 已經實作了 Salaried，因此可以當作 salary 的欄位值
	fmt.Println("Tim's salary is", tim.salary.getSalary())
	// output:
	// Tim's salary is 1200

	// anonymously nested interface
	ross := Woker2{
		firstName: "Ross",
		lastName:  "Geller",
		// 因為 Salary 實作了 Salaried，因此可以作為 Salaried 的欄位值
		Salaried: Salary{
			1000, 50, 50,
		},
	}

	// 由於 method 會被 promoted，因此可以直接呼叫 ross.getSalary() 的方法
	// 而不需要使用 ross.Salaried.getSalary()
	fmt.Println("Ross's salary is", ross.getSalary())
	// output:
	// Ross's salary is 1100

	// 匯出的欄位（Exported fields）
	// struct 中的欄位只有在欄位名稱以大寫命名時才會 export 出去，其他 package 中才能取用得到

	// struct 的比較（Struct comparison）
	// 當兩個 struct 的 type 和 field value 都相同時，兩個 struct 可以被視為相同
	// 但若在 struct 中有 field 的 type 是無法比較的話（例如，map），那麼這兩個 struct 將無法進行比較

	// 辨認 Struct Type 的名稱
	// 使用 reflect.TypeOf 和 reflect.ValueOf().Kind() 可以用來判斷該 struct 的 struct type 名稱，以及變數的實際 type
	u := struct {
		Name, Password string
	}{
		Name:     "Sammy the Shark",
		Password: "fisharegreat",
	}
	fmt.Println(reflect.TypeOf(u))
	fmt.Println(reflect.ValueOf(u).Kind())
	// output:
	// struct { Name string; Password string }
	// struct
}

// Human 有 name 且可以 Introduce，而 Saiyan 是 Human，因此它也有 Name 且可以 Introduce
type Human struct {
	Name string
}

func (h *Human) Introduce() {
	fmt.Printf("Hi, I'm %s\n", h.Name)
}

// 建立 Saiyan(賽亞人) struct，並將 Human embed 在內
// 意思是 Saiyan 是 Human 而不是 Saiyan「 有一個 」Human
type Saiyan struct {
	*Human
	Power int
}

// struct 中的欄位也可以放入 interface，以 Worker 這個 struct 來說，其中的 salary 欄位型別是 Salaried 這個 interface，也就是説 salary 這個欄位值，一定要有實作 Salarired 的方法，如此 salary 才會符合該 interface 的 type:
// 。Worker struct 中的 { salary Salaried} 表示 salary 要符合 Salaried interface type
// 。要符合該 interface type，表示 salary 要實作 Salaried interface 中所定義的 method
// 。在定義 ross 變數時，因為 Salary 這個 struct 已經實作了 Salaried，因此可以放到 salary 這個欄位中
type Salaried interface {
	getSalary() int
}

type Salary struct {
	basic, insurance, allowance int
}

func (s Salary) getSalary() int {
	return s.basic + s.insurance + s.allowance
}

type Worker struct {
	firstName, lastName string
	salary              Salaried // 只要 salary 實作了 Salaried，就可以 Salaried interface type
}

// 根據上述範例，若是有一個 struct 的欄位沒有填寫時 (anonymous fields)，interface 中所定義的方法也可以被 promoted：
// 。這裡定義一個 Woker2 struct 且使用了 Salaried 作為 anonymous field
// 。在對 Woker2 時，因為 Salary struct 有實作 Salaried，因此可以當作 Woker2 struct 中 Salaried 的值
// 。由於 promoted 這作用，可以直接使用 ross.getSalary() 方法，而不需要使用 ross.Salaried.getSalary()
type Woker2 struct {
	firstName, lastName string
	Salaried
}
