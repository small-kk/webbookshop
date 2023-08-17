package model

import (
	"fmt"
	"testing"
)

func TestOrderItem_CreateOrderItemTable(t *testing.T) {
	orderItem := &OrderItem{}
	err := orderItem.CreateOrderItemTable()
	fmt.Println(err)
}
