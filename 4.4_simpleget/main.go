package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
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
func main_4_8() {
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

// 4.9 multipart/form-data形式でファイルの送信
// `curl -F "name=Michel Jackson" -F "thumbnail=@photo.png" http://localhost:18888`
func main_4_9() {
	// bytes.Bufferを作成（メモリのバッファ、出力先）
	var buffer bytes.Buffer
	// multipart.Writerを作成（bufferへの書き込み器）
	writer := multipart.NewWriter(&buffer)

	// テキストフィールドを書き込む（"name"というパラメータ）
	writer.WriteField("name", "Michel Jackson")

	// ファイルフィールド用のライタを作成（"thumbnail"という名前でphoto.pngというファイル名）
	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.png")
	if err != nil {
		panic(err)
	}

	// ファイル内容を読み込む
	readFile, err := os.Open("photo.png")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	// ファイル内容をfileWriterにコピー（writerを通じてbufferに書き込まれる）
	io.Copy(fileWriter, readFile)

	// multipartデータの作成を完了（final boundaryを書き込む）
	writer.Close()

	// bufferの中身をPOSTリクエストとして送信
	// FormDataContentType()は "multipart/form-data; boundary=..." を返す
	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)

}

// 4.9.1 MIMEタイプを指定する
func main() {
	// bytes.Bufferを作成（メモリのバッファ、出力先）
	var buffer bytes.Buffer
	// multipart.Writerを作成（bufferへの書き込み器）
	writer := multipart.NewWriter(&buffer)

	// テキストフィールドを書き込む（"name"というパラメータ）
	writer.WriteField("name", "Michel Jackson")

	// MIMEタイプとContent-Dispositionヘッダーを設定するためのMIMEHeaderを作成
	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", "form-data; name=\"thumbnail\"; filename=\"photo.png\"")
	// partヘッダーを使用して、multipartリクエストに新しいパート（ファイルデータ）を作成
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.png")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	// ファイルの全コンテンツを、ファイルの書き込み用io.Writerにコピーします
	io.Copy(fileWriter, readFile)

	// multipartデータの作成を完了（final boundaryをバッファに書き込む）
	writer.Close()

	// bufferの中身をPOSTリクエストとして送信
	// FormDataContentType()は "multipart/form-data; boundary=..." を返す
	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

func main() {

}
