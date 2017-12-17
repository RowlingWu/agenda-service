package service
import (
  	//"encoding/json"
	"net/http"
	"github.com/RowlingWu/agenda-service/service/entities"
	"github.com/unrolled/render"
)

func registerHandler(formatter *render.Render) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    var user entities.User
    user.Name = r.Form["Name"][0]
    user.Passwd = r.Form["Passwd"][0]
    user.Email = r.Form["Email"][0]
    user.Tel = r.Form["Tel"][0]
  //  if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
	//		w.WriteHeader(http.StatusUnprocessableEntity)
	//		return
	//	}

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
