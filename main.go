package main

import (
	"fmt"
	"sync"
	"time"
)

// 存放肉品的 channel
var (
	meatsChan = make(chan string, 5) // 假設最多同時有5個肉品在處理
)

func main() {
	// 定義進貨肉品，由於map結構是無序的所以可以隨機取出肉品
	meats := map[int]string{}

	for i := 0; i < 10; i++ {
		// 放入牛肉
		meats[len(meats)] = "牛肉"

		// 放入豬肉
		if i < 7 {
			meats[len(meats)] = "豬肉"
		}

		// 放入雞肉
		if i < 5 {
			meats[len(meats)] = "雞肉"
		}
	}

	// 員工對應名稱
	workersName := []string{"A", "B", "C", "D", "E"}

	// 使用 WaitGroup 等待所有 goroutine 完成處理
	wg := &sync.WaitGroup{}
	wg.Add(len(workersName))

	// 創建 goroutine 模擬員工取肉和處理肉的過程
	for i := 0; i < len(workersName); i++ {
		mutex := &sync.Mutex{}
		go worker(workersName[i], wg, mutex)
	}

	// 把全部肉類放入 channel
	for _, meat := range meats {
		meatsChan <- meat
	}

	// 關閉 channel，等肉類全部放入 channel 後，就不允許再塞入 channel
	close(meatsChan)

	// 等待所有員工完成工作
	wg.Wait()
}

func worker(name string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	for meat := range meatsChan {
		mutex.Lock()
		fmt.Printf("%v 在 %s 取得%s\n", name, currentTime(), meat)
		processTime := 0
		switch meat {
		case "牛肉":
			processTime = 1
		case "豬肉":
			processTime = 2
		case "雞肉":
			processTime = 3
		}
		time.Sleep(time.Duration(processTime) * time.Second)
		fmt.Printf("%v 在 %s 處理完%s\n", name, currentTime(), meat)
		mutex.Unlock()
	}
}

func currentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
