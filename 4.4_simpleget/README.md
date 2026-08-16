# Real World HTTP - Go学習ログ

このディレクトリでは、Go言語を用いたHTTPクライアントの実装を段階的に学習しています。
セッションごとに `main` 関数を書き換え（または `main_x_x` 形式で保存し `main` から呼び出し）、機能を拡張していくスタイルで進めています。

## 学習ロードマップ

各ステップで実装した内容は以下の通りです。

### 4.4 基本的なGETリクエスト
- **目的**: GETメソッドの送信、レスポンスボディ、ステータスコード、ヘッダーの取得。
- **実装**: `main_4_4()`
- **確認内容**: ボディコンテンツ、`resp.Status`, `resp.StatusCode`, `resp.Header` の表示。

### 4.5 GETリクエストとクエリパラメータ
- **目的**: URLエンコードされたクエリパラメータの送信。
- **実装**: `main_4_5()`
- **対応するcurl**: `curl -G --data-urlencode "query=hello world" http://localhost:18888`

### 4.6 HEADメソッドの送信
- **目的**: ボディを除いたヘッダー情報のみの取得。
- **実装**: `main_4_6()`
- **対応するcurl**: `curl --head http://localhost:18888`

### 4.7 POSTメソッド (x-www-form-urlencoded)
- **目的**: フォーム形式でのデータ送信。
- **実装**: `main_4_7()`
- **対応するcurl**: `curl -d text-value http://localhost:18888`

### 4.8 POSTメソッド (任意コンテンツ)
- **目的**: バイナリデータやファイルの直接送信。
- **実装**: `main_4_8()`
- **対応するcurl**: `curl --data-binary @main.go -H "Content-Type: text/plain" http://localhost:18888`

### 4.9 multipart/form-data によるファイル送信
- **目的**: ファイルとテキストフィールドを同時に送信する形式の実装。
- **実装**: `main_4_9()`
- **対応するcurl**: `curl -F "name=Michel Jackson" -F "thumbnail=@photo.png" http://localhost:18888`

### 4.9.1 MIMEタイプの指定
- **目的**: 送信するパートごとに詳細なMIMEタイプ（`image/jpeg`など）を指定する方法。
- **実装**: `main_4_9_1()`

### 4.10 CookieJarによるクッキー管理
- **目的**: サーバーから送られたクッキーを保存し、次回以降のリクエストで自動送信する仕組み。
- **実装**: `main_4_10()`
- **ポイント**: `http.Client` に `cookiejar.New` で作成した Jar を設定。

### 4.11 プロキシの利用 (現在のメイン)
- **目的**: 特定のプロキシサーバーを経由してリクエストを送信する。
- **実装**: `main()`
- **対応するcurl**: `curl -x http://localhost:18888 http://github.com`
- **ポイント**: `http.Transport` の `Proxy` フィールドに `http.ProxyURL` を設定。

## 実行方法

現在の `main()` 関数を動作させるには以下のコマンドを実行します。

```bash
go run main.go
```

過去のステップを試したい場合は、`main()` 関数の中身を `main_x_x()` の呼び出しに書き換えて実行してください。
