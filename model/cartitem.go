package model

import "book_mall/repository"

// CartItem 购物项
type CartItem struct {
	ItemID int     `gorm:"PrimaryKey;AutoIncrement"`        //购物项的id
	Book   *Book   `gorm:"ForeignKey:BookID;References:ID"` //购物项图书信息
	BookID int     `gorm:"NotNull"`                         //将购物项中图书与books表中图书关联
	Count  int     `gorm:"NotNull"`                         //购物项图书数量
	Amount float64 `gorm:"NotNull"`                         //购物项图书金额小计
	CartID string  `gorm:"NotNull;Size:255"`                //当前购物项属于哪一个购物车
	Cart   Cart    `gorm:"ForeignKey:CartID;References:CartUUID"`
}

// GetAmount 获取购物项中图书金额小计
func (c *CartItem) GetAmount() float64 {
	//获取当前购物项中图书的价格
	price := c.Book.Price
	return float64(c.Count) * price
}

// CreateTable 建表
func (c *CartItem) CreateTable() error {
	err := repository.DB.Migrator().CreateTable(CartItem{})
	return err
}
