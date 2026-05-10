# テストエコーサーバー

## 起動
```shell
go run server.go
```

## curlコメンドの実行
```shell
curl --http1.0 http://localhost:18888/greeting

<html><body>hello<form action="/" method="POST" enctype="multipart/form-data"><input name="title"><input type="file"><input type="submit"></form></body></html>
```

## クエリパラメータを追加してリクエストする
```shell
curl --http1.0 --get --data-urlencode "search word" http://localhost:18888
```

## より詳細な情報を表示する
```shell
curl --http1.0 --get --data-urlencode "search word" http://localhost:18888 -v

* Uses proxy env variable NO_PROXY == '127.0.0.1,localhost'
* Host localhost:18888 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:18888...
* Connected to localhost (::1) port 18888
> GET /?search+word HTTP/1.0
> Host: localhost:18888
> User-Agent: curl/8.7.1
> Accept: */*
> 
* Request completely sent off
* HTTP 1.0, assume close after body
< HTTP/1.0 200 OK
< Date: Sat, 25 Apr 2026 04:47:19 GMT
< Content-Length: 160
< Content-Type: text/html; charset=utf-8
< 
```

