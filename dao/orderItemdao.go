package dao

import (
	"book_mall/model"
	"book_mall/repository"
)

// AddOrderItem 向数据库中添加订单项
func AddOrderItem(orderItem *model.OrderItem) error {
	err := repository.DB.Create(orderItem).Error
	return err
}

// GetOrderItemsByOrderID 根据订单号获取该订单所有订单项
func GetOrderItemsByOrderID(orderID string) ([]*model.OrderItem, error) {
	var orderItems []*model.OrderItem
	err := repository.DB.Where("order_id = ?", orderID).Find(&orderItems).Error
	return orderItems, err

}
