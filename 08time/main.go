package main

import (
	"fmt"
	"time"
)

func main() {
	// current time:
	// according to the local setting
	presentTime := time.Now()
	fmt.Println("presentTime: ", presentTime)
	// output:
	// presentTime: 2022-12-10 19:56:26.176486 +0800 CST m=+0.000113918

	// current time:
	// timestamp
	presentTimeStamp := time.Now().Unix()
	fmt.Println("presentTimeStamp: ", presentTimeStamp)
	// output:
	// presentTimeStamp:  1670673697

	// transfer timestamp to time
	transferTimeStampToTime := time.Unix(time.Now().Unix(), 0)
	fmt.Println("transferTimeStampToTime: ", transferTimeStampToTime)
	// output:
	// transferTimeStampToTime:  2022-12-10 20:04:47 +0800 CST

	// get current year, month, day
	y, m, d := time.Now().Date()
	fmt.Println("year: ", y, " month: ", m, " day: ", d)
	// output:
	// year:  2022  month:  December  day:  10
	// if I don't want to transfer December to 12,I can do :
	fmt.Println("year: ", y, " month: ", int(m), " day: ", d)
	// output:
	// year:  2022  month:  12  day:  10

	// get current hour, min, sec
	hou := time.Now().Hour()
	min := time.Now().Minute()
	sec := time.Now().Second()
	fmt.Println("hour: ", hou, " minute: ", min, " second: ", sec)
	// output:
	// hour:  20  minute:  10  second:  39

	// transfer int to date
	transferDate := time.Date(2022, 9, 9, 12, 12, 12, 0, time.UTC)
	fmt.Println("transferDate: ", transferDate)
	// output:
	// transferDate:  2022-09-09 12:12:12 +0000 UTC

	// format
	// 2006-01-05 15:04:05 is the rule you should type it if you want to get correct current format
	formatDate := time.Now().Format("2006-01-05 15:04:05")
	fmt.Println("formatDate: ", formatDate)
	// output:
	// ormatDate:  2022-12-03 20:16:03

	// advance to the next EX:時區轉換 =================================================================
	t := time.Now()
	var localLocation *time.Location
	var err error
	// 設定時區為UTC
	localLocation, err = time.LoadLocation("UTC")
	if err != nil {
		return
	}
	// local timezone transfer to utc timezone
	fmt.Println("UTC Timestamp: ", t.In(localLocation).Unix())
	// output:
	// UTC Timestamp:  1670740047

	fmt.Println("Local Time: ", t.In(localLocation))
	// output:
	// Local Time:  2022-12-11 06:30:36.501193 +0000 UTC

	fmt.Println("UTC Time: ", t)
	// output:
	// UTC Time:  2022-12-11 14:32:58.231025 +0800 CST m=+0.000387126

	// advance to the next EX:時間的計算 =================================================================
	// 計算 開始到結束的時間
	startTime := time.Now()                    // 開始時間
	time.Sleep(time.Duration(5) * time.Second) // 休息時間 也可以：time.Sleep(5 * time.Second)
	endTime := time.Since(startTime)           // 結束時間
	fmt.Println("endTime: ", endTime)
	// output:
	// endTime:  5.001053041s

	// 增加時間
	afterTime := startTime.Add(5 * time.Second) // 增加 5秒
	fmt.Println("startTime: ", startTime, "afterTime: ", afterTime)
	// output:
	// startTime:  2022-12-11 14:50:16.110385 +0800 CST m=+0.000323001 afterTime:  2022-12-11 14:50:21.110385 +0800 CST m=+5.000323001

	// 計算時間差距
	startTime2 := startTime.Add(-5 * time.Second) // 在 startTime 加上 負五秒 並賦值給 startTime2
	subDurations := startTime.Sub(startTime2)     // 計算 startTime 與 startTime2 的時間差
	fmt.Println("subDurations: ", subDurations)
	// output:
	// subDurations:  5s

	// advance to the next EX:定時器ticker =================================================================
	duration := time.Duration(2 * time.Second) // 設定 兩秒的時間間距
	// 設定 定時器 2秒 ticker
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	count := 0
	for {
		<-ticker.C // 此為channel 相關應用，之後講解
		if count == 10 {
			fmt.Println("stop do something")
			ticker.Stop()
			break
		}
		fmt.Println("do something")
		count++
	}
	// output:
	// do something
	// do something
	// do something
	// do something
	// do something
	// do something
	// do something
	// do something
	// do something
	// do something
	// stop do something

	// 補充
	// Reference: https://pjchender.dev/golang/defer-panic-recovery/
	// what is defer?
	// defer 可以用來在函式最終要回傳前被執行，此外 defer 有幾個原則：
	// 第一：defer中所使用到的參數是在執行到的時候就已經被帶入
	// language := "Go"
	// defer fmt.Print(language + "\n")
	// language = "Node.js"
	// fmt.Printf("Hello\n")
	// output:
	// Hello Go

	// 第二：當有多個defer時，執行順序為 後進先出 (last-input-first-output)
	// languageOne := "Go"
	// defer fmt.Printf("to " + languageOne + "\n")
	// languageTwo := "Node.js"
	// defer fmt.Printf("from " + languageTwo + "\n")
	// fmt.Printf("Hello\n")
	// output:
	// Hello 並沒有 defer 所以比defer先執行
	// from Node.js 在程式最後，最先執行
	// to Go 在程式開始，最後執行

	// 第三：defer 可以用在 函式回傳具名的變數，當函式回傳的值是具名的變數時，defer 可以去修改最終回傳出去的值：
	fmt.Println("deferEX: ", deferEX())
	// output:
	// deferEX:  2
	// 原本要輸出1後來在函式結束之前針對輸出的變數 i 加1，所以輸出 2
}

func deferEX() (i int) {
	defer func() { i++ }()
	return 1
}
