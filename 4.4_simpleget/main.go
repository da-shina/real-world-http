package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// 4.4 GETメソッドの送信とコンテンツ、ステータス、フィールドの表示
func main_4_4() {
	// GETメソッドの送信
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//  BODYコンテンツの表示
	log.Println(string(body))
	//  ステータスの表示
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
	//  フィールドの表示
	log.Println("Fields:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
}

// 4.5 GETメソッド・クエリーの送信
// `curl -G --data-urlencode "query=hello world" http://localhost:18888`
func main_4_5() {
	values := url.Values{
		"query": {"hello world"},
	}
	resp, err := http.Get("http://localhost:18888" + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Println(string(body))
}

// 4.6 HEADメソッドの送信
// `curl --head http://localhost:18888`
func main_4_6() {
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
}

// 4.7 POSTメソッドの送信（x-www-form-urlencoded形式）
// `curl -d text-value http://localhost:18888`
func main_4_7() {
	values := url.Values{
		"test": {"value"},
	}
	resp, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

// 4.8 POSTメソッドで任意コンテンツを送信
// `curl --data-binary @main.go -H "Content-Type: text/plain" http://localhost*18888`
func main() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)

}
