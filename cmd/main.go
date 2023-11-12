package main

import (
	router "RestAPIGolang/internal/api/routers"
	"RestAPIGolang/internal/database"
	"RestAPIGolang/internal/utils"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Flag for database fseeding
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()
	//Initialize database
	if err := database.Init(); err != nil {
		log.Fatal(err)
	}
	//seeding is allowed?
	if *seed {
		fmt.Println("seeding the database")
		utils.SeedAccounts()
	}
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
