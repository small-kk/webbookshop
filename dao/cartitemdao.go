package dao

import (
	"book_mall/model"
	"book_mall/repository"
)

// AddCartItem 插入购物项
func AddCartItem(cart *model.CartItem) error {
	err := repository.DB.Create(&cart).Error
	return err
}

// GetCartItemsByBookIDAndCartID 根据bookID来查询一个购物项
func GetCartItemsByBookIDAndCartID(bookID int, cartID string) (*model.CartItem, error) {
	cartItem := &model.CartItem{}
	err := repository.DB.Where("book_id = ? and cart_id = ?", bookID, cartID).Find(cartItem).Error
	return cartItem, err
}

// GetCartItemsByCartUUID 根据购物车的uuid获取uuid对应的所有购物项
func GetCartItemsByCartUUID(cartUUID string) ([]*model.CartItem, error) {
	var cartItems = []*model.CartItem{}
	err := repository.DB.Where("cart_id = ?", cartUUID).Find(&cartItems).Error
	for i := 0; i < len(cartItems); i++ {
		book, _ := GetBookByID(cartItems[i].BookID)
		cartItems[i].Book = book
	}
	return cartItems, err
}

// UpdatesCartItem 更新购物项信息
func UpdatesCartItem(cartItem *model.CartItem) error {
	err := repository.DB.Updates(cartItem).Error
	return err
}

// DeleteCartItemByCartUUID 根据购物车的UUID删除购物项
func DeleteCartItemByCartUUID(cartUUID string) error {
	err := repository.DB.Where("cart_id = ?", cartUUID).Delete(&model.CartItem{}).Error
	return err
}

// DeleteCartItemByCartItemID 根据购物项ID删除购物项
func DeleteCartItemByCartItemID(cartItemID string) error {
	err := repository.DB.Where("item_id=?", cartItemID).Delete(&model.CartItem{}).Error
	return err
}
