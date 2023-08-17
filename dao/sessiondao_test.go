package dao

import (
	"book_mall/model"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestAddSession(t *testing.T) {
	session := &model.Session{
		SessionID: uuid.NewString(),
		UserName:  "zhangsan",
		UserID:    1,
	}
	err := AddSession(session)
	fmt.Println(err)
}

func TestDeleteSession(t *testing.T) {
	sessID := "e8c2914c-31f9-418f-966d-ace1049da8a2"
	err := DeleteSession(sessID)
	fmt.Println(err)
}
