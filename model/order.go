package model

import (
	"book_mall/repository"
)

// Order 订单
type Order struct {
	OrderUUID   string  `gorm:"primaryKey"` //唯一订单号
	CreateTime  string  `gorm:"notNull"`    //订单生成时间
	TotalCount  int     `gorm:"notNull"`    //订单图书总数量
	TotalAmount float64 `gorm:"notNull"`    //订单图书总金额
	State       int8    `gorm:"notNull"`    //订单状态  0 表示未发货，1表示已发货，2表示交易完成
	UserID      int     `gorm:"notNUll"`    //订单所属用户
	User        *User   `gorm:"foreignKey:UserID;references:ID"`
}

func (o *Order) CreateOrderTable() error {
	err := repository.DB.Migrator().CreateTable(o)
	return err
}
