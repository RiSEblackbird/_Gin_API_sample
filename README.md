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
  - インポート
    - ``"_Gin_API_Sample/controller"``, ``"_Gin_API_Sample/middleware"``, ``"github.com/gin-gonic/gin"``, ``_ "github.com/go-sql-driver/mysql"``
  - デフォルトのミドルウェアでginルーターを作成する

- ``controller/book.go``
  - インポート
    - ``"_Gin_API_Sample/model"``, ``"_Gin_API_Sample/service"``, ``"net/http"``, ``"strconv"``, ``"github.com/gin-gonic/gin"``
  - CRUD操作の関数を定義
    - ``BookAdd()``, ``BookList()``, ``BookUpdate()``, ``BookDelete()``

- ``middleware/bookMiddleware.go``
  - インポート
    - ``"log"``, ``"time"``, ``"github.com/gin-gonic/gin"``, ``"go.uber.org/zap"``
  - ``zap``パッケージを使用したリクエストのロギング
    - ``RecordUaAndTime()``

- ``model/book.go``
  - ``xorm``の書式でbookの構造体を定義

- ``service/init.go``
  - データベースへの接続とテーブルの初期化

- ``service/book.go``

## Tips

### import文について

- Import declarations : https://golang.org/ref/spec#Import_declarations
  - "_"(unserscore)
    - パッケージの副作用 (初期化) のみを目的としてインポートする場合に空白識別子を前方に付して記述する

      ~~~go
      import _ "lib/math"
      ~~~

### ポインタ関連

- [Pointers - A Tour of Go](https://go-tour-jp.appspot.com/moretypes/1)
- [Go's Declaration Syntax - The Go Blog](https://blog.golang.org/declaration-syntax)
- [Goで学ぶポインタとアドレス - Qiita](https://qiita.com/Sekky0905/items/447efa04a95e3fec217f)

### ``go.mod``内の``indirect``コメント

- [Using Go Modules - The Go Blog](https://blog.golang.org/using-go-modules)
  - ある依存関係がそのモジュールによって直接使用されるのではなく、他のモジュールの依存関係によって間接的にのみ使用されることを示す

### 型(types)

- [Types - The Go Programming Language Specification](https://golang.org/ref/spec#Types)
  - [#Struct types](https://golang.org/ref/spec#Struct_types)

### 構造体(struct)

- [Goを学びたての人が誤解しがちなtypeと構造体について #golang - Qiita](https://qiita.com/tenntenn/items/45c568d43e950292bc31)

## 階層

~~~txt
_Gin_API_Sample
├── controller
│   └── book.go
├── middleware
│   └── bookMiddleware.go
├── .gitignore               // standard gitignore file
├── go.mod
├── go.sum
├── main.go
└── README.md                // simple readme file

### ├── │ └──
~~~

## Error

### ``  ``

~~~error
~~~

- ````
- 要因&対処
  - 
