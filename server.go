// mainパッケージの定義
package main

// 必要なパッケージのインポート
import (
	"fmt"               // 書式化されたI/Oを扱うためのパッケージ
	"log"               // ロギング機能を提供するパッケージ
	"net/http"          // HTTPクライアントとサーバーの実装を提供するパッケージ
	"net/http/httputil" // HTTPリクエストとレスポンスのダンプ機能を提供するパッケージ
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello<form action=\"/\" method=\"POST\" enctype=\"multipart/form-data\"><input name=\"title\"><input type=\"file\"><input type=\"submit\"></form></body></html>\n")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
