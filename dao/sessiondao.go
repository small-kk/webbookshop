package dao

import (
	"book_mall/model"
	"book_mall/repository"
)

// AddSession 向数据库中添加session
func AddSession(sess *model.Session) error {
	err := repository.DB.Create(&sess).Error
	return err
}

// DeleteSession 删除数据库中的session
func DeleteSession(sessID string) error {
	err := repository.DB.Where("session_id=?", sessID).Delete(&model.Session{}).Error
	return err
}

// GetSessionBySessionID 根据sessionID从数据库中查询
func GetSessionBySessionID(sessID string) (*model.Session, error) {
	session := &model.Session{}
	err := repository.DB.Where("session_id = ?", sessID).Find(&session).Error
	return session, err
}
