package cmd

import (
	"encoding/json"
	"testing"

	"github.com/RowlingWu/agenda-service/entity"
)

// test from  API mocker server
var testAllUsersURL = "https://private-59ac4-test15398.apiary-mock.com/v1/user/getkey?username=root&password=pass"
var testRegisterURL = ""
var testLoginURL = ""
var testLogoutURL = ""
var testIsLoginURL = ""

// users got from API mocker server should be the same as testAllUsers
var testAllUsers = []entity.User{{1, "usertest", "secret", "xxx@qq.com", "136"}, {2, "xxx", "sec", "abc@mail.com", "199"}}

func TestQuery(t *testing.T) {
	b, err := query(testAllUsersURL)
	logError(err, t)

	var users []entity.User
	err = json.Unmarshal(b, &users)
	logError(err, t)

	for i, u := range users {
		if u != testAllUsers[i] {
			t.Fatal("getUsers not equal to setUsers", "getUsers:", u, "setUsers:", testAllUsers[i])
		}
	}
}

func TestRegister(t *testing.T) {
	err := register(&entity.User{}, testRegisterURL)
	logError(err, t)
}

func TestLogin(t *testing.T) {
	err := login("", "", testLoginURL)
	logError(err, t)
}

func TestLogout(t *testing.T) {
	err := logout("", testLogoutURL)
	logError(err, t)
}

func TestIsLogin(t *testing.T) {
	_, err := isLogin(testIsLoginURL)
	logError(err, t)
}

func logError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
