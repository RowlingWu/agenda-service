package service

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func NewServer() *negroni.Negroni {

		formatter := render.New()

		n := negroni.Classic()
		mx := mux.NewRouter()

		initApiRoutes(mx, formatter)

	  //n.Use(&CheckIsLoginHandler{})
		n.UseHandler(mx)
		return n
}

func initApiRoutes(mx *mux.Router, formatter *render.Render) {

	// Group User
	mx.HandleFunc("/v1/user/login", isLoginHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/user/login", loginHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/user/logout", logoutHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/user/self", deleteHandler(formatter)).Methods("DELETE")

	// Group Users
	mx.HandleFunc("/v1/users", listAllUsersHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/users", registerHandler(formatter)).Methods("POST")
}
