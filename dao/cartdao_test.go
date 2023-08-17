package dao

import (
	"book_mall/model"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestAddCart(t *testing.T) {
	//购物车session_id
	cartSessionId := uuid.NewString()
	item1 := &model.CartItem{
		Book: &model.Book{
			ID:     1,
			Title:  "解忧杂货店",
			Author: "东野圭吾",
			Price:  27,
			Sales:  100,
		},
		Count:  2,
		CartID: cartSessionId,
	}
	item1.Amount = item1.GetAmount()
	item1.BookID = item1.Book.ID

	item2 := &model.CartItem{
		Book: &model.Book{
			ID:     2,
			Title:  "中国哲学史",
			Author: "冯友兰",
			Price:  44,
			Sales:  101,
		},
		Count:  3,
		CartID: cartSessionId,
	}
	item2.Amount = item2.GetAmount()
	item2.BookID = item2.Book.ID

	cart := &model.Cart{
		CartUUID:  cartSessionId,
		CartItems: []*model.CartItem{item1, item2},
		UserID:    1,
	}
	cart.TotalAmount = cart.GetTotalAmount()
	cart.TotalCount = cart.GetTotalCount()
	err := AddCart(cart)
	fmt.Println(err)

}

func TestGetCartByUserID(t *testing.T) {
	cart, _ := GetCartByUserID(2)
	fmt.Println(cart)
	fmt.Println(cart.CartUUID == "")
	fmt.Println(cart == nil)
	for _, item := range cart.CartItems {
		fmt.Println(item)
	}
}

func TestUpdatesCart(t *testing.T) {
	cart := &model.Cart{
		CartUUID:    "2f887cc3-4a0b-4398-8fb3-73db395f9087",
		TotalAmount: 54,
		TotalCount:  2,
		UserID:      2,
	}
	err := UpdatesCart(cart)
	fmt.Println(err)
}

func TestDeleteCartByCartUUID(t *testing.T) {
	err := DeleteCartByCartUUID("2f887cc3-4a0b-4398-8fb3-73db395f9087")
	fmt.Println(err)
}
