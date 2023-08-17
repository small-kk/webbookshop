package controller

import (
	"book_mall/dao"
	"book_mall/utils"
	"html/template"
	"net/http"
)

func ToIndex(w http.ResponseWriter, r *http.Request) {

	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//调用bookdao中带有分页的图书函数
	page, _ := dao.GetPageBooks(pageNo)

	//判断是否已经登录
	session, ok := utils.IsLogin(r)
	if ok {
		page.IsLogin = true
		page.UserName = session.UserName
	}

	//解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	//渲染模板
	t.Execute(w, page)
}
