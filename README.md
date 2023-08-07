* 題目說明:
想像你擁有一間肉品加工廠，工廠只加工三種肉類，牛、豬、雞。今天有一堆肉品進 貨，共有牛肉 10 份、豬肉 7 份、雞肉 5 份。
你有五個員工，代號為 A,B,C,D,E，每一個員工處理肉的速度都一樣，牛肉需要 1 秒，豬 肉需要 2 秒，雞肉需要 3 秒，五個員工彼此的工作互不干涉，他們都能夠獨立處理肉品，且 每個人一次只能處理一塊肉，每塊肉只能被一個人處理，人取肉後不得放回。
你的目標是把這次進貨的牛肉 10 份、豬肉 7 份、雞肉 5 份，交由五位員工全部處理完 成，每一位員工都是隨機取肉，並請你忽略取肉的時間，唯一會耗時的只有【處理肉】。
請模擬上述情境，用 concurrency 的概念撰寫一支程式，將所有肉品處理完畢。

2. Golang 請使用 goroutine 配合 channel 撰寫。必定要有一個 channel 用來放所有肉品。
3. 請詳細印出取肉、處理肉的過程，並附註時間
例如:
A 在 2022-01-01 15:00:00 取得牛肉
B 在 2022-01-01 15:00:01 取得豬肉
C 在 2022-01-01 15:00:01 取得雞肉
A 在 2022-01-01 15:00:01 處理完牛肉 A 在 2022-01-01 15:00:01 取得豬肉
A 在 2022-01-01 15:00:03 處理完豬肉 B 在 2022-01-01 15:00:03 處理完豬肉 ......(以下省略)

4. 請確認沒有人【同時拿了同一塊肉】。
5. 不能修改進貨的肉量和員工數量。
6. 請妥善架構你的程式，並加上人能看懂的註解。
7. 本題目沒有提供標準輸入或輸出，請自行揣摩，沒有一定的寫法。
8. 只要程式沒有 error，都歡迎寄回。