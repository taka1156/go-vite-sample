package dbconnect

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func isErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@tcp(db:3306)/sample?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true")
	if err != nil {
		fmt.Print("データベース接続に失敗しました。")
	}
	db.LogMode(true)
	return db
}
