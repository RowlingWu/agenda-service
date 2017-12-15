package service
import (
  "encoding/json"
	"net/http"
	"agenda-service/service/entities"
	"github.com/unrolled/render"
)

func registerHandler(formatter *render.Render) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    var user entities.User
    json.NewDecoder(r.Body).Decode(&user)

    entities.UserServ.MyAdd(&user)
    formatter.JSON(w, http.StatusCreated,struct{}{})
    return
  }
}

func deleteHandler(formatter *render.Render) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    user := isLogin(r)
    entities.UserServ.MyDelete(user)
    logoutHandler(formatter)(w,r)
  }
}

func listAllUsersHandler(formatter *render.Render) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    formatter.JSON(w, http.StatusOK, entities.UserServ.ListAll())
  }
}
