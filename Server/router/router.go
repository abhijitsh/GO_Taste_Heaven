package router

import (
	"github.com/abhijitsh/go_restapi/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	//GetAllUsers and AddUser should be a controller
	r.HandleFunc("/api/user", middleware.GetAllUsers).Methods("GET")
	r.HandleFunc("/api/user", middleware.AddUser).Methods("POST")
	r.HandleFunc("/api/signup", middleware.SignUp).Methods("POST")
	r.HandleFunc("/api/login", middleware.Login).Methods("POST")
	r.HandleFunc("/api/address", middleware.PostAddress).Methods("POST")
	r.HandleFunc("/api/address", middleware.GetAllAddresses).Methods("GET")
	r.HandleFunc("/api/confirmation", middleware.GetConfirmation).Methods("GET")
	r.HandleFunc("/api/order", middleware.PostOrder).Methods("POST")
	r.HandleFunc("/api/order", middleware.GetOrder).Methods("GET")
	r.HandleFunc("/api/profile", middleware.PostProfile).Methods("POST")
	r.HandleFunc("/api/profile", middleware.GetProfile).Methods("GET")
	r.HandleFunc("/api/image", middleware.UpdateIMG).Methods("POST")
	r.HandleFunc("/api/cart", middleware.PostCart).Methods("POST")
	r.HandleFunc("/api/cart", middleware.GetCart).Methods("GET")
	//TODO:Create route to delete user as well
	return r
}
