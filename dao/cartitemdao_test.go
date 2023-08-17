package dao

import (
	"book_mall/model"
	"fmt"
	"testing"
)

func TestGetCartItemsByBookIDAndCartID(t *testing.T) {
	cartItem, _ := GetCartItemsByBookIDAndCartID(3, "b63eef9e-46f5-4d82-827c-9370c644a7fb")
	fmt.Println(cartItem)
}

func TestGetCartItemsByCartUUID(t *testing.T) {
	cartItems, _ := GetCartItemsByCartUUID("b63eef9e-46f5-4d82-827c-9370c644a7fb")
	for _, cartItem := range cartItems {
		fmt.Println(*cartItem)
	}
}

func TestUpdatesCartItem(t *testing.T) {
	cartItem := &model.CartItem{
		ItemID: 7,
		BookID: 1,
		Count:  2,
		Amount: 54,
		CartID: "2f887cc3-4a0b-4398-8fb3-73db395f9087",
	}
	err := UpdatesCartItem(cartItem)
	fmt.Println(err)
}
