package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Name      string
	Pwd       string
	tableName string
}

func (u *User) TableName() string {
	// custom table name, this is default
	return u.tableName
}
func main() {
	db, err := gorm.Open("mysql", "root:@/wallet_development?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	db.AutoMigrate(&User{tableName: "user1"})
	db.AutoMigrate(&User{tableName: "user2"})

	db.Create(&User{tableName: "user1", Name: "n1", Pwd: "p1"})
	db.Create(&User{tableName: "user2", Name: "n2", Pwd: "p2"})
}
