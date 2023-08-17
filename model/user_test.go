package model

import "testing"

// 执行测试函数，创建用户表
func TestUser_CreateTable(t *testing.T) {
	u := User{}
	u.CreateTable()
}
