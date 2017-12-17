package cmd

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/RowlingWu/agenda-service/entity"
)

func query(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte(""), err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return []byte(""), errors.New(res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), err
	}
	return body, nil
}

func register(u *entity.User, rurl string) error {
	resp, err := http.PostForm(rurl, url.Values{"Name": {u.Name}, "Passwd": {u.Passwd}, "Email": {u.Email}, "Tel": {u.Tel}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(resp.Status + "\n" + string(body))
	}

	return nil
}

func login(name string, pwd string, rurl string) error {
	resp, err := http.PostForm(rurl, url.Values{"Name": {name}, "Passwd": {pwd}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(resp.Status + "\n" + string(body))
	}
	return nil
}

func logout(name string, rurl string) error {
	resp, err := http.PostForm(rurl, url.Values{"Name": {name}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(resp.Status + "\n" + string(body))
	}
	return nil
}

func isLogin(rurl string) ([]byte, error) {
	resp, err := http.Get(rurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 401 {
		return []byte("User not logged in."), nil
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status + "\n" + string(body))
	}

	return body, nil
}

func deleteUser(rurl string) (bool, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", rurl, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true, nil
	}
	if resp.StatusCode == 401 {
		return false, nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	return false, errors.New(resp.Status + "\n" + string(body))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
