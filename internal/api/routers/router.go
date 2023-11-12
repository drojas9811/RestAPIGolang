package router

import (
	"RestAPIGolang/internal/api/handlers"
	"RestAPIGolang/internal/api/middlewares"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/login", middlewares.MakeHTTPHandleFunc(handlers.Login)).Methods("POST")
	router.HandleFunc("/account", middlewares.MakeHTTPHandleFunc(handlers.CreateAccount)).Methods("POST")
	router.HandleFunc("/account", middlewares.MakeHTTPHandleFunc(handlers.GetAccount)).Methods("GET")
	router.HandleFunc("/account/{id}", middlewares.WithJWTAuth(middlewares.MakeHTTPHandleFunc(handlers.GetAccountByID))).Methods("GET")
	router.HandleFunc("/account/{id}", middlewares.WithJWTAuth(middlewares.MakeHTTPHandleFunc(handlers.DeleteAccountByID))).Methods("DELETE")
	router.HandleFunc("/transfer", middlewares.MakeHTTPHandleFunc(handlers.Transfer)).Methods("POST")

	return router
}
