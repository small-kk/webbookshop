package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("验证用户是否能登录成功", testCheckUserNameAndPassword)
	t.Run("验证用户是否能注册成功", testCheckUserNameAndRegister)
}

// testCheckUserNameAndPassword 验证用户是否能登录成功
func testCheckUserNameAndPassword(T *testing.T) {
	b, err := CheckUserNameAndPassword("zhaoliu", "123123")
	fmt.Println(b, err)
}

// testCheckUserNameAndRegister 验证用户是否能注册成功
func testCheckUserNameAndRegister(t *testing.T) {
	b, err := CheckUserNameAndRegister("zhangsan03", "10010", "wangwu@sina.com")
	fmt.Println(b, err)
}
