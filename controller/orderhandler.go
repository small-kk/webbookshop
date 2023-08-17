package controller

import (
	"book_mall/dao"
	"book_mall/model"
	"book_mall/utils"
	"github.com/google/uuid"
	"html/template"
	"net/http"
	"time"
)

// Checkout 去结账
func Checkout(w http.ResponseWriter, r *http.Request) {
	//获取session
	session, _ := utils.IsLogin(r)
	//获取用户id
	userID := session.UserID
	//获取购物车
	cart, _ := dao.GetCartByUserID(userID)
	//获取用户信息
	user, _ := dao.GetUserByUserID(userID)

	//创建订单
	order := &model.Order{
		OrderUUID:   uuid.NewString(),
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      userID,
		User:        user,
	}
	//将订单保存在数据库
	dao.AddOrder(order)
	//保存订单项
	//获取购物车中的购物项
	for _, cartItem := range cart.CartItems {
		//创建订单项
		orderItem := &model.OrderItem{
			Count:   cartItem.Count,
			Amount:  cartItem.Amount,
			Title:   cartItem.Book.Title,
			Author:  cartItem.Book.Author,
			Price:   cartItem.Book.Price,
			ImgPath: cartItem.Book.ImgPath,
			OrderID: order.OrderUUID,
		}
		//将订单项保存在数据库中
		dao.AddOrderItem(orderItem)
		//更新当前订单项图书的销量和库存
		book := cartItem.Book
		book.Sales = book.Sales + orderItem.Count
		book.Stock = book.Stock - orderItem.Count
		dao.UpdatesBook(book)
	}
	//清空购物车
	dao.DeleteCartByCartUUID(cart.CartUUID)
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w, order)
}

// GetOrders 获取所有订单
func GetOrders(w http.ResponseWriter, r *http.Request) {
	//获取所有订单
	orders, _ := dao.GetOrders()
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}

// GetOrderInfo 获取订单详情
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	//获取订单对应的订单项
	orderID := r.FormValue("orderId")
	orderItems, _ := dao.GetOrderItemsByOrderID(orderID)
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w, orderItems)
}

// GetMyOrders 获取我的订单
func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.IsLogin(r)
	userID := session.UserID
	orders, _ := dao.GetMyOrder(userID)
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w, struct {
		UserName string
		Orders   []*model.Order
	}{
		UserName: session.UserName,
		Orders:   orders,
	})
}

// SendOrder 发货
func SendOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderId")
	//发货
	dao.UpdateOrderState(orderID, 1)
	//再次获取所有订单信息
	GetOrders(w, r)
}

// TakeOrder 收货
func TakeOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderId")
	dao.UpdateOrderState(orderID, 2)
	GetMyOrders(w, r)
}
