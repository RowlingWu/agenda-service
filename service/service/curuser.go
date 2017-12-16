package service

import (
  "encoding/json"
	"net/http"
	"github.com/RowlingWu/agenda-service/service/entities"
	"github.com/satori/go.uuid"
	"github.com/unrolled/render"
)

var (
  nonLoginAPI = map[string]bool {
    "v1/user/login": true,
    "v1/users":      true,
  }
)

type CheckIsLoginHandler struct{}

func (h *CheckIsLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  _, isNonLogin := nonLoginAPI[r.URL.Path]
	if !isNonLogin && isLogin(r) == nil {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		next(w, r)
	}
}

func isLogin(r *http.Request) *entities.User {
  usernameCookie, err := r.Cookie("username")
  if err != nil {
    return nil
  }
  username := usernameCookie.Value
  curuser := entities.CurServ.CurQuery(username)
  if curuser == nil {
    return nil
  }
  user := entities.UserServ.MyQuery(username)
  if user == nil {
    return nil
  }
  return user
}

func isLoginHandler(formatter *render.Render) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if user := isLogin(r); user != nil {
      formatter.JSON(w,http.StatusOK,user)
    } else {
      formatter.JSON(w,http.StatusUnauthorized,struct{}{})
    }
  }
}

func loginHandler(formatter *render.Render) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    var invalid struct {
      MSG string `json:"msg"`
    }
    if user:=isLogin(r); user != nil {
      invalid.MSG = "Please log out " + user.Name +" first"
      formatter.JSON(w,http.StatusBadRequest,invalid)
      return
    }

    var requestBody struct {
      Name string `json:Name`
      Passwd string `json:Passwd`
    }
    if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

    user := entities.UserServ.MyQuery(requestBody.Name)
    if user == nil || user.Passwd != requestBody.Passwd {
      invalid.MSG = "Wrong username or password"
      formatter.JSON(w,http.StatusUnauthorized,invalid)
    }

    var username string
    for {
      username = uuid.NewV4().String()
      if entities.CurServ.CurQuery(username) == nil {
        break
      }
    }

    curname := entities.CurUser{
      Username: requestBody.Name,
    }
    entities.CurServ.MyAdd(&curname)
    formatter.JSON(w, http.StatusOK, struct {
      MSG string `json:"msg"`
    }{
      MSG: "Log in successfully",
    })
  }
}

func logoutHandler(formatter *render.Render) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    usernameCookie,_:=r.Cookie("username")
    username := usernameCookie.Value
    entities.CurServ.MyDelete(username)
    w.WriteHeader(http.StatusOK)
  }
}
