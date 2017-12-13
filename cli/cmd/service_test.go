package cmd

import (
	"encoding/json"
	"testing"

	"github.com/RowlingWu/agenda-service/entity"
)

// users got from API mocker server should be the same as testAllUsers
var testAllUsers = []entity.User{{1, "usertest", "secret", "xxx@qq.com", "136"}, {2, "xxx", "sec", "abc@mail.com", "199"}}

func TestQuery(t *testing.T) {
	b := query(testAllUsersURL)
	var users []entity.User
	err := json.Unmarshal(b, &users)
	logError(err, t)
	for i, u := range users {
		if u != testAllUsers[i] {
			t.Fatal("getUsers not equal to setUsers", "getUsers:", u, "setUsers:", testAllUsers[i])
		}
	}
}

func logError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
