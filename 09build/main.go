package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi")

	// 這個章節要介紹 go 的指令 build 及 env
	// go 可以 藉由 build指令 去產生一個執行檔，也可以藉由 env 這個指令 查看 go 的環境變數
	// 可以在 terminal中 輸入 go env
	// 這邊我們要用的環境變數為 GOOS：go 目前的操作系統
	// GOOS="darwin" => macOS
	// GOOS="windows" => windows
	// GOOS="linux" => linux

	// 再來若想要產生不同操作系統的執行檔，可以在 terminal 中輸入：
	// GOOS="darwin" go build => macOS 的執行檔
	// output:
	// buildIntro

	// GOOS="windows" go build  => windows 的執行檔
	// output:
	// buildIntro.exe

	// GOOS="linux" go build => linux 的執行檔
	// output:
	// buildIntro

	// 補充:
	// Refrence:
	// 1.https://pjchender.dev/golang/env/
	// 2.https://medium.com/%E4%BC%81%E9%B5%9D%E4%B9%9F%E6%87%82%E7%A8%8B%E5%BC%8F%E8%A8%AD%E8%A8%88/golang-goroot-gopath-go-modules-%E4%B8%89%E8%80%85%E7%9A%84%E9%97%9C%E4%BF%82%E4%BB%8B%E7%B4%B9-d17481d7a655
	// 針對 go env 的某些參數做講解：

	// GOPATH: =================================================
	// 專門存放第三方套件以供我們程式碼的需要，其中主要包含三個資料夾：src、pkg、bin。
	// 在沒有使用 GO Modules前，GOPATH 會是所有工作環境的根目錄，而GOPATH的一個很大的缺點，那就是你相關的第三方套件只要不是官方程式庫，都需要放置在GOPATH/src的路徑下才可以使用，所以之後才有了 Go Modules。

	// $GOPATH/src:
	// 非官方的第三方套件所放置的位置，不過後來go1.11之後 GO Modules 出現了，所以就沒再使用了
	// $GOPATH/bin:
	// 主要放的是 執行 go install時，GO編譯後的執行檔。
	// 一般來說會建議把這個資料夾加到系統上的 global $PATH 變數，如此直接在終端機中執行這些編譯後的執行檔。
	// $GOPATH/pkg:
	// 這里主要是保存一些編譯前的物件，用來簡短編譯所需的時間，
	// 一般來說，開發者並不需要進來這裡。若在編譯時碰到錯誤，可以安心地把這個資料夾刪除，Go 會自己再建立新的。

	// GOROOT: =================================================
	// GOROOT 裡面放的是 Go 語言本身自己要用的東西，或是內建的函式庫，像是 Go 編譯成執行檔時所用的工具。
	// 預設的情況下，$GOROOT 的路徑通常會是 /use/local/go，而 $GOPATH 會是 $HOME/go。
}
