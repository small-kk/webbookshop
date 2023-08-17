package model

import (
	"fmt"
	"testing"
)

func TestCreateOrderTable(t *testing.T) {
	order := &Order{}
	err := order.CreateOrderTable()
	fmt.Println(err)
}
