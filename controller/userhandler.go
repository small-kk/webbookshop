package controller

import (
	"book_mall/dao"
	"book_mall/model"
	"book_mall/repository"
	"book_mall/utils"
	"fmt"
	"github.com/google/uuid"
	"html/template"
	"net/http"
)

// Logout 处理用户注销
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取cookie
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {
		//获取cookie的value值
		cookieValue := cookie.Value
		//删除数据库中对应的session
		dao.DeleteSession(cookieValue)
		//设置cookie失效
		cookie.MaxAge = -1
		//将修改之后的cookie发送个浏览器
		http.SetCookie(w, cookie)
	}

	//去首页
	GetPageBooksByPrice(w, r)
}

// Login 处理用户登录函数
func Login(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登录
	_, ok := utils.IsLogin(r)
	if ok {
		//已经登录
		GetPageBooksByPrice(w, r)
	} else {
		//获取用户名和密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		//调用userdao中验证用户名和密码方法
		user, err := dao.CheckUserNameAndPassword(username, password)
		if err != nil {
			fmt.Println("CheckUserNameAndPassword err:", err)
		}
		if user.ID != 0 {
			//用户名和密码正确
			//创建Session
			sess := &model.Session{
				SessionID: uuid.NewString(),
				UserName:  user.UserName,
				UserID:    user.ID,
			}
			//将sess保存在数据库中
			dao.AddSession(sess)

			//创建一个Cookie，使其与Session相关联
			http.SetCookie(w, &http.Cookie{
				Name:     "session_id",
				Value:    sess.SessionID,
				HttpOnly: true,
			})
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user)
		} else {
			//用户名或密码不正确
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "用户名或密码不正确！")
		}
	}

}

// Register 注册处理
func Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	repwd := r.PostFormValue("repwd")
	email := r.PostFormValue("email")

	//两次密码不相等，这一块可以不用，js帮忙处理了
	if password != repwd {
		t := template.Must(template.ParseFiles("views/pages/regist.html"))
		t.Execute(w, nil)
	} else {
		//检查并注册
		b, err := dao.CheckUserNameAndRegister(username, password, email)
		if err != nil {
			fmt.Println("CheckUserNameAndRegister err:", err)
		}

		if b {
			//b为true，表示注册成功
			t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
			t.Execute(w, nil)
		} else {
			//b为false，表示用户名已存在，注册失败
			t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
			t.Execute(w, "用户名已存在！")
		}
	}
}

// CheckUserName 通过发送Ajax请求验证用户名是否可用
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user := model.User{}
	repository.DB.Where("user_name = ?", username).Find(&user)
	if user.ID != 0 {
		//表示用户名存在
		w.Write([]byte("用户名已存在！"))
	} else {
		//表示用户名可用
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}
