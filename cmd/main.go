package main

import (
	router "RestAPIGolang/internal/api/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Initialize
	if err :=Init(); err!=nil{ log.Fatal("Web server (HTTPS): ", err)}
	//Routes definition
	router := router.InitRouter()
	//Port definition
	port := 3000
	//Start server
	log.Println("JSON API server running on port: ", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Fatal("Web server (HTTPS): ", err)
	}

}
