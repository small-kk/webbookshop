package main

import (
	"book_mall/controller"
	"net/http"
)

func main() {
	//处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))

	//去首页
	http.HandleFunc("/index", controller.ToIndex)

	//去登陆
	http.HandleFunc("/login", controller.Login)

	//去注销
	http.HandleFunc("/logout", controller.Logout)

	//去注册
	http.HandleFunc("/regist", controller.Register)

	//通过Ajax请求验证用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)

	//获取每一页图书信息
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	//http.HandleFunc("/getPageBooks", controller.GetBooks)
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)

	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)

	//去添加或更新图书页面
	http.HandleFunc("/toAddOrUpdateBookPage", controller.ToAddOrUpdateBookPage)

	//去添加或更新图书
	http.HandleFunc("/addOrUpdatesBook", controller.AddOrUpdatesBook)

	//添加图书到购物车
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)

	//获取购物车信息
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)

	//清空购物车
	http.HandleFunc("/deleteCart", controller.DeleteCart)

	//删除购物项
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)

	//更新购物项
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)

	//去结账
	http.HandleFunc("/checkout", controller.Checkout)

	//获取所有订单
	http.HandleFunc("/getOrders", controller.GetOrders)

	//获取订单详情
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)

	//获取我的订单
	http.HandleFunc("/getMyOrder", controller.GetMyOrders)

	//发货
	http.HandleFunc("/sendOrder", controller.SendOrder)

	//确认收获
	http.HandleFunc("/takeOrder", controller.TakeOrder)

	http.ListenAndServe(":9090", nil)

}

//func main() {
//	r := gin.Default()
//	r.Static("/css", "./views/static/css")
//	r.Static("/img", "./views/static/img")
//	r.Static("/script", "./views/static/script")
//
//	r.LoadHTMLGlob("./views/pages/**/*.html")
//
//	r.GET("/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "index.html", "")
//	})
//
//	pagesGroup := r.Group("/pages")
//	{
//		cartGroup := pagesGroup.Group("/cart")
//		{
//			cartGroup.GET("/cart.html", func(c *gin.Context) {
//				c.HTML(http.StatusOK, "cart.html", nil)
//			})
//
//			cartGroup.GET("/checkout.html", func(c *gin.Context) {
//				c.HTML(http.StatusOK, "checkout.html", nil)
//			})
//		}
//
//		managerGroup := pagesGroup.Group("/manager")
//		{
//			managerGroup.GET("/book_edit.html", func(c *gin.Context) {
//				c.HTML(http.StatusOK, "book_edit.html", nil)
//			})
//
//			managerGroup.GET("/book_manager.html", func(c *gin.Context) {
//				c.HTML(http.StatusOK, "book_manager.html", nil)
//			})
//
//			managerGroup.GET("/manager.html", func(c *gin.Context) {
//				c.HTML(http.StatusOK, "manager.html", nil)
//			})
//		}
//
//		orderGroup := pagesGroup.Group("/order")
//		{
//			orderGroup.GET("/order.html", func(c *gin.Context) {
//				c.HTML(http.StatusOK, "order.html", nil)
//			})
//
//			orderGroup.GET("/order_info.html", func(c *gin.Context) {
//				c.HTML(http.StatusOK, "order_info.html", nil)
//			})
//
//			orderGroup.GET("/order_manager.html", func(c *gin.Context) {
//				c.HTML(http.StatusOK, "order_manager.html", nil)
//			})
//		}
//
//		userGroup := pagesGroup.Group("/user")
//		{
//			userGroup.GET("/login.html", func(c *gin.Context) {
//				c.HTML(http.StatusOK, "login.html", nil)
//			})
//
//			userGroup.GET("/login_success.html", func(c *gin.Context) {
//				c.HTML(http.StatusOK, "login_success.html", nil)
//			})
//
//			userGroup.GET("/regist.html", func(c *gin.Context) {
//				c.HTML(http.StatusOK, "regist.html", nil)
//			})
//
//			userGroup.GET("/regist_success.html", func(c *gin.Context) {
//				c.HTML(http.StatusOK, "regist_success.html", nil)
//			})
//		}
//	}
//
//	r.Run(":9090")
//}
