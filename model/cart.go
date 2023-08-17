package model

import (
	"book_mall/repository"
)

// Cart 购物车
type Cart struct {
	CartUUID    string      `gorm:"PrimaryKey"` //购物车id
	CartItems   []*CartItem `gorm:"-"`          //购物车中所有的购物项
	TotalCount  int         `gorm:"NotNull"`    //购物车中图书总数量
	TotalAmount float64     `gorm:"NotNull"`    //购物车中总金额
	UserID      int         `gorm:"NotNull"`    //当前购物车所属用户
	User        *User       `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
}

// GetTotalCount 获取购物车中图书总数量
func (c *Cart) GetTotalCount() int {
	var totalCount int
	for _, v := range c.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

// GetTotalAmount 获取购物车中图书总金额
func (c *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range c.CartItems {
		totalAmount += v.Amount
	}
	return totalAmount
}

// CreateTable 建表
func (c *Cart) CreateTable() error {
	err := repository.DB.Migrator().CreateTable(Cart{})
	return err
}
