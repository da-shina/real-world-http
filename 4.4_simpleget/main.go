package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

// 4.4 GETメソッドの送信とコンテンツ、ステータス、フィールドの送信
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
func main() {
	values := url.Values{
		"query": {"hello world"},
	}
	resp, _ := http.Get("http://localhost:18888" + "?" + values.Encode())
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Println(string(body))
}
