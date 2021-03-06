package entities_test

import (
	"entities"
	"fmt"
	"go/build"
	"runtime"
	"testing"
)

func init() {
	if runtime.GOOS == "windows" {
		entities.Initial(build.Default.GOPATH + "\\tmp\\test_agenda.db")
	} else {
		entities.Initial("/tmp/test_agenda.db")
	}
}

func TestCurUser(t *testing.T) {
	t.Log("test curuser MyAdd")
	fmt.Println("check cur add")
	username := "TEST"
	entities.CurServ.MyAdd(&entities.CurUser{
		Username: username,
	})
	if entities.CurServ.CurQuery(username) == nil {
		t.Fatalf("cannot find user added just now, add failed")
	}
	t.Log("delete the test user")
	entities.CurServ.MyDelete(username)
}
