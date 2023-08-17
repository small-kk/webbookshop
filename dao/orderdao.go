package dao

import (
	"book_mall/model"
	"book_mall/repository"
)

// AddOrder 向数据库中插入订单
func AddOrder(order *model.Order) error {
	err := repository.DB.Create(order).Error
	return err
}

// GetOrders 获取数据库中所有的订单
func GetOrders() ([]*model.Order, error) {
	var orders []*model.Order
	err := repository.DB.Find(&orders).Error
	return orders, err
}

// GetMyOrder 获取我的订单
func GetMyOrder(userID int) ([]*model.Order, error) {
	var orders []*model.Order
	err := repository.DB.Where("user_id =?", userID).Find(&orders).Error
	return orders, err
}

// UpdateOrderState 更新订单状态
func UpdateOrderState(orderUUID string, state int) error {
	err := repository.DB.Model(model.Order{}).Where("order_uuid = ?", orderUUID).Update("state", state).Error
	return err
}
