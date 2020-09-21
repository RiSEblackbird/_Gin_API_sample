# _Gin_API_Sample

(This repository is my own self-study document
)

## リポジトリの目的

- ``Gin``の基礎学習
  - とりあえず組んで動かす
- 学習記録/自分用資料まとめ

## 主なリファレンス

- メインチュートリアル
  - ***[ginを最速でマスターしよう - Qiita](https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a)***

### **Go**について

- GitHub : https://github.com/golang/go

- Goコマンド
  - 原文 : https://golang.org/cmd/go/
  - 日訳 : https://godoc.org/github.com/gophersjp/go/src/cmd/go

- 文法 (The Go Programming Language Specification)
  - 原文 : https://golang.org/ref/spec

### **Gin**について

- GitHub : https://github.com/gin-gonic/gin

## 工程(参考先より適宜変更)

### パッケージ管理

- [GOMODULE--Goのパッケージ管理 - Qiita](https://qiita.com/Syoitu/items/f221b52231703cebe8ff)
- ``$ go mod init _Gin_API_Sample``

### パッケージ導入

- ``$ go get -u github.com/gin-gonic/gin``
- ``$ go get github.com/go-sql-driver/mysql``
  - **go-sql-driver/mysql** : https://github.com/go-sql-driver/mysql
- ``$ go get xorm.io/xorm``
  - **xorm / xorm** : https://gitea.com/xorm/xorm
  - [旧] go-xorm/xorm : https://github.com/go-xorm/xorm
    - Go向けのORMツール
- ``$ go get go.uber.org/zap``
  - **uber-go/zap** : https://github.com/uber-go/zap
    - 高速で動作するロギングツール

### APIの作成([ginを最速でマスターしよう - Qiita](https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a#gin%E3%81%A7%E7%B0%A1%E5%8D%98%E3%81%AArest%E9%A2%A8%E3%81%AEapi%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC%E3%82%92%E4%BD%9C%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%88%E3%81%86))

- ``main.go``
  - 各パッケージのインポート
  - デフォルトのミドルウェアでginルーターを作成する

## Tips

### import文について

 - Import declarations : https://golang.org/ref/spec#Import_declarations
   - "_"(unserscore)
      - パッケージの副作用 (初期化) のみを目的としてインポートする場合に空白識別子を前方に付して記述する
        ~~~go
        import _ "lib/math"
        ~~~

## 階層

~~~txt
_Gin_API_Sample
├── .gitignore               // standard gitignore file
├── go.mod
├── go.sum
├── main.go
└── README.md                // simple readme file

### ├── │ └─
~~~

## Error

### ``  ``

~~~error
~~~

- ````
- 要因&対処
  - 
