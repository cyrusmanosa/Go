package main

import (
	"io"
	"log"
	"net/http"
)

// http.ListenAndServe開始建立服務，並監聽我們設定的埠，內有兩個參數分別是address
// 成功執行後在127.0.0.1:8080即可看見顯示Hello Web的網頁了！

// --HandleFunc
// net/http package中 HandleFunc 方法，將func test() 方法與 "/" 這個 routing 進行綁定，讓 server 知道當進來的 traffic 的 routing 為 / 時要執行 test 方法
// 加入 index.html 並綁定相關的 request handler => http.HandleFunc("/index", test)

// --ListenAndServe
// 開始運行伺服器 => http.ListenAndServe( "address（URL or POST）" , (handle) )
// 可用作Check Server

// --http.ResponseWriter, r *http.Request（資料驗證）
// request handler 中回傳 html

// 在response內回傳html格式的文字
func test1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	str := `<!DOCTYPE html>
	<html>
		<head>
			<title>首頁</title>
		</head>
		<body>
			<h1>首頁</h1>
			<p>啦啦啦啦啦啦啦</p>
		</body>
	</html>`
	io.WriteString(w, str)
}

type HtmlData struct {
	Title   string
	Content string
}

func main() {
	http.HandleFunc("/", test1)              // url 路徑以 "/" 開頭的 request 都會被 test1 function 處理．
	err := http.ListenAndServe(":8081", nil) // address: 本例設定了要監聽的port為8080。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} // if ERROR
}
