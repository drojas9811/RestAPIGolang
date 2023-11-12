package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/anthdm/gobank/internal/api/helpers"
	router "github.com/anthdm/gobank/internal/api/routers"
	"github.com/anthdm/gobank/internal/database"
)

func main(){
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
		helpers.SeedAccounts()
	}
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


