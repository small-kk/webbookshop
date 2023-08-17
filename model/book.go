package model

import "book_mall/repository"

type Book struct {
	ID      int     `gorm:"primaryKey"`
	Title   string  `gorm:"notNull"`
	Author  string  `gorm:"notNull"`
	Price   float64 `gorm:"notNull"`
	Sales   int     `gorm:"notNull"`
	Stock   int     `gorm:"notNull"`
	ImgPath string  `gorm:"notNull"`
}

// CreateBookTable 利用gorm自动创建一个数据库表
func (b Book) CreateBookTable() {
	repository.DB.Migrator().CreateTable(b)
}
