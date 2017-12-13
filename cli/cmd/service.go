package cmd

import (
	"io/ioutil"
	"net/http"
)

// get all users from API mocker server
var testAllUsersURL = "https://private-59ac4-test15398.apiary-mock.com/v1/user/getkey?username=root&password=pass"

func query(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return body
}
