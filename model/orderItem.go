package model

import "book_mall/repository"

// OrderItem 订单项
type OrderItem struct {
	OrderItemID int     `gorm:"primaryKey;AutoIncrement"` //订单项的Id，自增
	Count       int     `gorm:"notNull"`                  //订单项中图书数量
	Amount      float64 `gorm:"notNull"`                  //订单项中图书金额小计
	Title       string  `gorm:"notNull"`                  //订单项图书书名
	Author      string  `gorm:"notNull"`                  //订单项图书作者
	Price       float64 `gorm:"notNull"`                  //订单项中图书的金额
	ImgPath     string  `gorm:"notNull"`                  //订单项中图书封面
	OrderID     string  `gorm:"notNull;Size:255"`         //订单项所属的订单
	Order       *Order  `gorm:"foreignKey:OrderID;references:OrderUUID"`
}

func (o *OrderItem) CreateOrderItemTable() error {
	err := repository.DB.Migrator().CreateTable(o)
	return err
}
