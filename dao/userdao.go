package dao

import (
	"book_mall/model"
	"book_mall/repository"
	"errors"
	"fmt"
)

// CheckUserNameAndPassword 查询用户是否存在
func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	var user *model.User
	err := repository.DB.Where("user_name=? and pass_word=?", username, password).Find(&user).Error
	return user, err
}

// CheckUserNameAndRegister 检查用户是否存在，若存在提醒请更换用户名，如果用户不存在，则注册，将用户信息保存在数据库中
func CheckUserNameAndRegister(username string, password string, email string) (bool, error) {
	var user model.User
	repository.DB.Where("user_name=?", username).Find(&user)
	fmt.Println(user)

	if user.ID != 0 {
		return false, errors.New("用户已存在，请更换用户名！")
	} else {
		u := model.User{
			UserName: username,
			PassWord: password,
			Email:    email,
		}
		repository.DB.Create(&u)
		return true, nil
	}
}

// GetUserByUserID 根据用户id获取用户信息
func GetUserByUserID(userID int) (*model.User, error) {
	user := &model.User{}
	err := repository.DB.Where("id= ?", userID).Find(user).Error
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
