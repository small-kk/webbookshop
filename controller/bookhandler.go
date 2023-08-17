package controller

import (
	"book_mall/dao"
	"book_mall/model"
	"book_mall/utils"
	"html/template"
	"net/http"
	"strconv"
)

// GetBooks 获取所有图书
//func GetBooks(w http.ResponseWriter, r *http.Request) {
//	//调用bookdao中GetBooks函数
//	books, _ := dao.GetBooks()
//	//解析模板文件
//	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
//	t.Execute(w, books)
//}

// GetPageBooks 获取带分页的图书信息
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//调用bookdao中带有分页的图书函数
	page, _ := dao.GetPageBooks(pageNo)
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}

// GetPageBooksByPrice 获取带分页和价格范围的图书
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}

	//获取价格范围
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")

	var page *model.Page

	if minPrice == "" && maxPrice == "" {
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		//将价格范围设置到page中
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	////获取cookie
	//cookie, _ := r.Cookie("session_id")
	//if cookie != nil {
	//	//查询session_id是否存在
	//	session, _ := dao.GetSessionBySessionID(cookie.Value)
	//	if session.UserID > 0 {
	//		//已经登录
	//		page.IsLogin = true
	//		page.UserName = session.UserName
	//	}
	//}
	session, ok := utils.IsLogin(r)
	if ok {
		page.IsLogin = true
		page.UserName = session.UserName
	}

	//解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}

// DeleteBook 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//bookId := r.URL.Query().Get("bookId")
	bookId := r.FormValue("bookId")
	iBookId, _ := strconv.Atoi(bookId)
	book := &model.Book{
		ID: iBookId,
	}
	dao.DeleteBook(book)
	GetPageBooks(w, r)
}

// ToAddOrUpdateBookPage 去添加或更新图书的页面
func ToAddOrUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	iBookId, _ := strconv.Atoi(bookId)
	book, _ := dao.GetBookByID(iBookId)
	if book.ID != 0 {
		//去更新页面
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	} else {
		//去添加页面
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, nil)
	}
}

// AddOrUpdatesBook 添加图书或更新图书处理器
func AddOrUpdatesBook(w http.ResponseWriter, r *http.Request) {

	//判断是添加图书还是更新图书
	//添加图书，r中没有bookId
	//更新图书，r中有bookId

	id := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	price := r.PostFormValue("price")
	author := r.PostFormValue("author")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	iid, _ := strconv.Atoi(id)
	fprice, _ := strconv.ParseFloat(price, 64)
	isales, _ := strconv.Atoi(sales)
	istock, _ := strconv.Atoi(stock)

	book := &model.Book{
		Title:   title,
		Author:  author,
		Price:   fprice,
		Sales:   isales,
		Stock:   istock,
		ImgPath: "static/img/default.jpg",
	}
	//不管是添加图书还是更新图书，最后都需要去图书管理页面，将图书信息显示出来
	if iid != 0 {
		//iid，即为bookId不为0，表示在更新图书
		book.ID = iid
		dao.UpdatesBook(book)

	} else {
		//iid，即bookId为0，表示在添加图书
		dao.AddBook(book)
	}
	//从数据库中获取最新图书信息
	GetPageBooks(w, r)

}
