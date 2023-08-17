package dao

import (
	"book_mall/model"
	"fmt"
	"testing"
)

func TestAddOrderItem(t *testing.T) {
	orderItem := &model.OrderItem{
		Count:   2,
		Amount:  100,
		Title:   "童话",
		Author:  "安徒生",
		Price:   50.00,
		ImgPath: "static/img/default.jpg",
		OrderID: "5dd6b6e1-c5b5-497f-a0e8-c911fa1272c1",
	}
	err := AddOrderItem(orderItem)
	fmt.Println(err)
}

func TestGetOrderItemsByOrderID(t *testing.T) {
	orderItems, _ := GetOrderItemsByOrderID("e3f86138-ec68-4625-af45-01e6812a5e0c")
	for _, orderItem := range orderItems {
		fmt.Println(orderItem)
	}
}
