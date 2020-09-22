package service

import (
	"_Gin_API_Sample/model"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	DsName := "gintutorial:password@/ginsample"
	err := errors.New("")
	// 新しくDBのエンジンを定義
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	// Sync2のテーブル構造体を定義
	DbEngine.Sync2(new(model.Book))
	fmt.Println("init data base ok")
}
