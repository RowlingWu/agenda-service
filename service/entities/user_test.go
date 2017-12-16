package entities_test

import (
	"github.com/RowlingWu/agenda-service/service/entities"
	"testing"
)

func TestUser(t *testing.T) {
	t.Log("test MyAdd")
        username := "arcus"
	u := &entities.User{
		Name: username,
		Passwd: "123456",
		Email:    "123@123.com",
		Tel:    "1234567890",
	}
	entities.UserServ.MyAdd(u)
	if entities.UserServ.MyQuery(username) == nil {
		t.Fatalf("cannot find user added just now, add failed")
	}
	t.Log("delete the test user")
	entities.UserServ.MyDelete(u)
}
