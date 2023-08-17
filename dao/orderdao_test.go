package dao

import (
	"book_mall/model"
	"fmt"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestAddOrder(t *testing.T) {
	order := &model.Order{
		OrderUUID:   uuid.NewString(),
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  2,
		TotalAmount: 100,
		State:       0,
		UserID:      1,
	}
	err := AddOrder(order)
	fmt.Println(err)
}

func TestGerOrders(t *testing.T) {
	orders, _ := GetOrders()
	for _, v := range orders {
		fmt.Println(v)
	}
	fmt.Println(orders)
}

func TestGetMyOrder(t *testing.T) {
	orders, _ := GetMyOrder(1)
	for _, order := range orders {
		fmt.Println(order)
	}
}

func TestUpdateOrderState(t *testing.T) {
	err := UpdateOrderState("e3f86138-ec68-4625-af45-01e6812a5e0c", 1)
	fmt.Println(err)
}
