package main

import (
	"fmt"
	"log"
	"net/http"
	router "github.com/anthdm/gobank/internal/api/routers"
	"github.com/anthdm/gobank/internal/database"
)

func main(){

	//Initialize database
	database.GetDB()
	//Routes definition
	router:=router.InitRouter()
	//Port definition
	port:= 8080
	//Start server
	log.Println("JSON API server running on port: ", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
        log.Fatal("Web server (HTTPS): ", err)
    }
}


