package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler 會處理傳送到 "/" 路徑的請求。
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 將 "Hello, World!" 字串寫入回應中。
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	// 將根路徑 ("/") 的所有請求都交給 helloHandler 處理。
	http.HandleFunc("/", helloHandler)

	// 在終端機顯示伺服器正在啟動的訊息。
	fmt.Println("伺服器即將在 http://localhost:5500 上啟動")

	// 啟動伺服器並監聽 5500 連接埠。
	// 如果啟動失敗（例如連接埠已被佔用），程式會記錄錯誤並退出。
	if err := http.ListenAndServe(":5500", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
