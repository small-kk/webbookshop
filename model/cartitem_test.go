package model

import (
	"fmt"
	"testing"
)

func TestCartItem_CreateTable(t *testing.T) {
	c := CartItem{}
	err := c.CreateTable()
	fmt.Println(err)
}
