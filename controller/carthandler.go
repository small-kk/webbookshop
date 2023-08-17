package controller

import (
	"book_mall/dao"
	"book_mall/model"
	"book_mall/utils"
	"fmt"
	"github.com/google/uuid"
	"html/template"
	"net/http"
	"strconv"
)

// AddBook2Cart 添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登录
	session, ok := utils.IsLogin(r)
	if ok {
		//已经登录
		//获取添加的图书ID
		bookID := r.PostFormValue("bookId")
		//根据图书ID获取图书信息
		ibookID, _ := strconv.Atoi(bookID)
		book, _ := dao.GetBookByID(ibookID)

		//获取用户ID
		userID := session.UserID
		//判断数据库中是否有当前用户的购物车
		cart, _ := dao.GetCartByUserID(userID)
		if cart.CartUUID != "" {
			//当前用户已经有购物车
			//判断该用户购物车中是否存在有当前这本图书
			cartItem, _ := dao.GetCartItemsByBookIDAndCartID(ibookID, cart.CartUUID)
			if cartItem.BookID == ibookID {
				//表示已经存在这本书
				cartItem.Count++
				cartItem.Book = book
				cartItem.Amount = cartItem.GetAmount()
				//更新购物项信息
				dao.UpdatesCartItem(cartItem)
				//更新购物车中被修改购物项的信息
				for i := 0; i < len(cart.CartItems); i++ {
					if cart.CartItems[i].BookID == ibookID {
						cart.CartItems[i] = cartItem
						break
					}
				}
				cart.TotalCount = cart.GetTotalCount()
				cart.TotalAmount = cart.GetTotalAmount()
				//在数据库中更新购物车信息
				dao.UpdatesCart(cart)

			} else {
				//表示没有这本书
				//创建一个新的购物项
				newCartItem := &model.CartItem{
					Book:   book,
					BookID: ibookID,
					Count:  1,
					Amount: book.Price,
					CartID: cart.CartUUID,
				}
				//更新已存在的购物车中信息
				cart.CartItems = append(cart.CartItems, newCartItem)
				cart.TotalCount = cart.GetTotalCount()
				cart.TotalAmount = cart.GetTotalAmount()

				//将新创建的购物项添加到数据库中
				dao.AddCartItem(newCartItem)
				//更新购物车信息
				dao.UpdatesCart(cart)
			}
		} else {
			//当前用户没有购物车
			//创建购物车
			newCart := &model.Cart{
				CartUUID: uuid.NewString(),
				UserID:   userID,
			}
			//创建购物项
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				BookID: ibookID,
				Count:  1,
				Amount: book.Price,
				CartID: newCart.CartUUID,
			}
			cartItems = append(cartItems, cartItem)

			//补全购物车信息
			newCart.CartItems = cartItems
			newCart.TotalAmount = newCart.GetTotalAmount()
			newCart.TotalCount = newCart.GetTotalCount()

			//将购物车newCart保存到数据库中
			dao.AddCart(newCart)
		}
		w.Write([]byte(fmt.Sprintf("您刚刚将%s添加到购物车中", book.Title)))
	} else {
		//没有登录
		w.Write([]byte("请先登录！"))
	}
}

// GetCartInfo 根据用户id获取购物车信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.IsLogin(r)
	//获取用户id
	userID := session.UserID
	//根据用户id从数据库中获取该用户的购物车
	cart, _ := dao.GetCartByUserID(userID)
	//将用户信息保存在购物车中
	user, _ := dao.GetUserByUserID(userID)
	cart.User = user
	if cart.CartUUID != "" {
		//该用户有购物车
		//解析购物车模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, cart)
	} else {
		//用户没有购物车
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, nil)
	}
}

// DeleteCart 清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	//获取购物车id
	cartID := r.FormValue("cartId")
	//清空购物车
	dao.DeleteCartByCartUUID(cartID)
	//再次查询购物车信息
	GetCartInfo(w, r)
}

// DeleteCartItem 删除购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	//更新购物车信息
	session, _ := utils.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)

	icartItemID, _ := strconv.Atoi(cartItemID)
	for index, cartItem := range cart.CartItems {
		if cartItem.ItemID == icartItemID {
			cart.CartItems = append(cart.CartItems[:index], cart.CartItems[index+1:]...)
			//删除购物项
			dao.DeleteCartItemByCartItemID(cartItemID)
			break
		}
	}
	cart.TotalCount = cart.GetTotalCount()
	cart.TotalAmount = cart.GetTotalAmount()
	dao.UpdatesCart(cart)
	//获取最新购物车信息
	GetCartInfo(w, r)
}

// UpdateCartItem 更新购物项
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	icartItemID, _ := strconv.Atoi(cartItemID)
	bookCount := r.FormValue("bookCount")
	ibookCount, _ := strconv.Atoi(bookCount)
	//更新购物车信息
	session, _ := utils.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	for _, cartItem := range cart.CartItems {
		if cartItem.ItemID == icartItemID {
			//cart.CartItems[index].Count = ibookCount
			//cart.CartItems[index].Amount = cart.CartItems[index].GetAmount()
			////更新购物项信息
			//dao.UpdatesCartItem(cart.CartItems[index])
			cartItem.Count = ibookCount
			cartItem.Amount = cartItem.GetAmount()
			dao.UpdatesCartItem(cartItem)
			break
		}
	}
	cart.TotalCount = cart.GetTotalCount()
	cart.TotalAmount = cart.GetTotalAmount()
	dao.UpdatesCart(cart)
	//获取最新购物车信息
	GetCartInfo(w, r)
}
