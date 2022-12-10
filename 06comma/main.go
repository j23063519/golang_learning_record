package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcom := "Welcome to user input"
	fmt.Println(welcom)

	// os.Stdin
	// standard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok / err ok
	// after inputing something you can get what you input and error
	// if you don't want get error you can use _ to ignore it
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is %T \n", input)

	// 補充：
	// what is the bufio?
	// what is the os?
	// 在介紹這兩個套件之前，要先知道 io package 這個套件，
	// 因爲 上述兩個套件都是 基於 io package 製作而成的

	// what is io?
	// io這個套件事處理輸入與輸出等行為，分別定義了io.Reader及io.Writer接口
	// 分別用來抽象化輸入及輸出的操作
	// 若想要詳細知道io這個套件可以查看 README.md 的第二個參考網址

	// what is bufio
	// bufio這個套件基於io套件實現了可以藉由緩存實現讀取與寫入等行為，
	// 通過緩存提高效率，通過把文件放入緩存裡面，避免了直接讀取寫入文件，進而提高速度
	// 相同的寫入也是如此。
	// 詳細教學，可以查看 README.md 的第三個參考網址

	// what is os
	// os這個套件提供不依賴於 操作系統方法 的函式。
	// 也可以針對檔案進行寫入讀取等等行為。
	// 詳細教學，可以查看 README.md 的第四個參考網址
}
