package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// helloHandler 會處理傳送到 "/" 路徑的請求。
func helloHandler(c *gin.Context) {
	// 將 "Hello, World!" 字串寫入回應中。
	c.String(http.StatusOK, "Hello, World!")
}

func main() {
	// 建立一個預設的 Gin router，它已經包含了 Logger 和 Recovery middleware。
	r := gin.Default()

	// 將根路徑 ("/") 的所有請求都交給 helloHandler 處理。
	r.GET("/", helloHandler)

	// 在終端機顯示類似 Flask 的啟動訊息
	fmt.Println(" * Serving Flask app 'main'")
	fmt.Println(" * Debug mode: on")
	fmt.Println("WARNING: This is a development server. Do not use it in a production deployment. Use a production WSGI server instead.")
	fmt.Println(" * Running on http://127.0.0.1:5500")
	fmt.Println("Press CTRL+C to quit")
	fmt.Println("PR Test")
	fmt.Println("PR Test")

	// 啟動伺服器並監聽 5500 連接埠。
	r.Run("0.0.0.0:5500")
}
