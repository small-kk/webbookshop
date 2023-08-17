package utils

import (
	"book_mall/dao"
	"book_mall/model"
	"net/http"
)

// IsLogin 判断用户是否已经登录
func IsLogin(r *http.Request) (*model.Session, bool) {
	//根据cookie的name获取cookie
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {
		//获取cookie的value
		cookieValue := cookie.Value
		//根据cookie的value去数据库sessions查找对应的session
		session, _ := dao.GetSessionBySessionID(cookieValue)
		if session.UserID > 0 {
			return session, true
		}
	}
	return nil, false
}
