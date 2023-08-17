package model

import "book_mall/repository"

type Session struct {
	SessionID string `gorm:"PrimaryKey"`
	UserName  string `gorm:"NotNull;Unique"`
	UserID    int    `gorm:"NotNull;Unique"`
	User      User   `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
}

func (s *Session) CreateTable() {
	repository.DB.Migrator().CreateTable(Session{})
}
