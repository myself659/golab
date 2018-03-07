package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Name      string
	Pwd       string
	Alias     string
	tableName string
}

func (u *User) TableName() string {
	// custom table name, this is default
	return "user3"
}
func main() {
	db, err := gorm.Open("mysql", "root:jbigc@/wallet_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&User{})
	if !db.HasTable("user1") {
		db.CreateTable(&User{tableName: "user1"})
	}
	if !db.HasTable("user2") {
		db.CreateTable(&User{tableName: "user2"})
	}
	user3 := User{Name: "n3", Pwd: "p3"}
	db.Create(&user3)
	db.Create(&User{Name: "n1", Pwd: "p1"})
	db.Create(&User{Name: "n2", Pwd: "p2"})
}
