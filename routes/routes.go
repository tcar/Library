package routes

import (
	"net/http"

	"github.com/tcar/Library/controllers"
	"github.com/tcar/Library/utils/jwt"
)

func CreateRoutes(mux *http.ServeMux, uc *controllers.UserController) {
	//user routes
	mux.HandleFunc("/register", uc.Register)
	mux.HandleFunc("/login", uc.Login)
	mux.HandleFunc("/logout", uc.Logout)
	mux.HandleFunc("/secure", jwt.Authorization(uc.Secure))
}
