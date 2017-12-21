package cmd

import (
	"encoding/json"
	"testing"

	"github.com/RowlingWu/agenda-service/entity"
)

// test from  API mocker server
var testAllUsersURL = "https://private-7ef3f4-agenda30.apiary-mock.com/v1/user/users"
var testRegisterURL = "https://private-7ef3f4-agenda30.apiary-mock.com/v1/user/users"
var testLoginURL = "https://private-7ef3f4-agenda30.apiary-mock.com/v1/user/login"
var testLogoutURL = "https://private-7ef3f4-agenda30.apiary-mock.com/v1/user/logout"
var testIsLoginURL = "https://private-7ef3f4-agenda30.apiary-mock.com/v1/user/login/admin"
var testDeleteURL = "https://private-7ef3f4-agenda30.apiary-mock.com/v1/user/self"

// users got from API mocker server should be the same as testAllUsers
var testAllUsers = []entity.User{{
	1,
	"agenda",
	"123456",
	"agendaserver@111.com",
	"123456789"}, {
	1,
	"Tom",
	"123456",
	"agendaserver@222.com",
	"0987654321"}}

func TestQuery(t *testing.T) {
	b, err := query(testAllUsersURL)
	logError(err, t)
	t.Log("Success in accessing ListAllUsers' URL(mock).")

	var users []entity.User
	err = json.Unmarshal(b, &users)
	logError(err, t)

	for i, u := range users {
		if u != testAllUsers[i] {
			t.Fatal("getUsers not equal to setUsers", "getUsers:", u, "setUsers:", testAllUsers[i])
		}
	}

	t.Log("Users got from mock URL are correct")
}

func TestRegister(t *testing.T) {
	err := register(&entity.User{}, testRegisterURL)
	logError(err, t)
	t.Log("Success in accessing Register's URL(mock)\nStatus is 201 Created")
}

func TestLogin(t *testing.T) {
	err := login("xxx", "xxx", testLoginURL)
	logError(err, t)
	t.Log("Success in accessing Login's URL(mock)\nStatus is 200 OK")
}

func TestLogout(t *testing.T) {
	err := logout("", testLogoutURL)
	logError(err, t)
	t.Log("Success in accessing Logout's URL(mock).\nStatus is 200 OK")
}

func TestIsLogin(t *testing.T) {
	_, err := isLogin(testIsLoginURL)
	logError(err, t)
	t.Log("Success in accessing LoginStatus' URL(mock)\nReturned user's log status")
}

func TestDeleteUser(t *testing.T) {
	_, err := deleteUser(testDeleteURL)
	logError(err, t)
	t.Log("Success in accessing DeleteUser's URL(mock)\nReturned the result of DeleteUser")
}

func logError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
