package model

import (
	"book_mall/repository"
	"fmt"
)

// User 用户类型
type User struct {
	ID       int    `gorm:"primaryKey"`
	UserName string `gorm:"notNull;unique"`
	PassWord string `gorm:"notNull"`
	Email    string
}

// CreateTable 创建数据库表
func (u User) CreateTable() {
	err := repository.DB.Migrator().CreateTable(u)
	if err != nil {
		fmt.Println("create mysql table failed,err:", err)
		return
	}
	fmt.Println("add user success!")
}

// AddUser 添加用户
func (u *User) AddUser() {
	err := repository.DB.Create(u).Error
	if err != nil {
		fmt.Println("insert data failed,err:", err)
		return
	}
}
