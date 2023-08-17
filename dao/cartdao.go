package dao

import (
	"book_mall/model"
	"book_mall/repository"
)

// AddCart 向购物车表中插入购物车
func AddCart(cart *model.Cart) error {
	err := repository.DB.Create(&cart).Error
	for _, item := range cart.CartItems {
		AddCartItem(item)
	}
	return err
}

// GetCartByUserID 根据用户的ID从数据库中查询对应的购物车
func GetCartByUserID(userID int) (*model.Cart, error) {
	cart := &model.Cart{}
	//var cart *model.Cart
	err := repository.DB.Where("user_id=?", userID).Find(&cart).Error
	if err != nil {
		return nil, err
	}
	//获取当前购物车中所有的购物项
	cartItems, _ := GetCartItemsByCartUUID(cart.CartUUID)
	cart.CartItems = cartItems
	return cart, err
}

// UpdatesCart 更新购物车信息
func UpdatesCart(cart *model.Cart) error {
	err := repository.DB.Updates(cart).Error
	return err
}

// 根据购物车的UUID删除购物车
func DeleteCartByCartUUID(cartUUID string) error {
	err := DeleteCartItemByCartUUID(cartUUID)
	if err != nil {
		return err
	}
	err = repository.DB.Where("cart_uuid = ?", cartUUID).Delete(&model.Cart{}).Error
	if err != nil {
		return err
	}
	return nil
}
