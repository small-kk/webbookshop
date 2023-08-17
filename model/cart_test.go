package model

import (
	"fmt"
	"testing"
)

func TestCart_CreateTable(t *testing.T) {
	c := Cart{}
	err := c.CreateTable()
	fmt.Println(err)
}
