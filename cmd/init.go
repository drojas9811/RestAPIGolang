package main

import (
	"RestAPIGolang/internal/aws"
	"RestAPIGolang/internal/database"
	"RestAPIGolang/internal/utils"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func Init() error{
	log.Printf("%s", "LOG_INFO::INIT::MSG::App is initializing.")
	if err:=initializeDataBase(); err!=nil{ return errors.New("there was an issue in initializing the database")}
	log.Printf("%s", "LOG_INFO::INIT::MSG::Database is ready.")
	if err:=sedingDataBase(); err!=nil{ return errors.New("there was an issue in initializing the database")}
	log.Printf("%s", "LOG_INFO::INIT::MSG::Database has been seeded.")
	if err:=settingEnv(); err!=nil{ return errors.New("there was an issue in initializing the database")}
	log.Printf("%s", "LOG_INFO::INIT::MSG::Enviroment variable has been setted.")
	return nil
}

func sedingDataBase() error {
	//Flag for database fseeding
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	//seeding is allowed?
	if *seed {
		log.Printf("%s","LOG_INFO::INIT::sedingDataBase::MSG::seeding the database.")
		utils.SeedAccounts()
	}
	return nil
}

func initializeDataBase() error{
	//Initialize database
	if err := database.Init(); err != nil {
		errorString:= fmt.Sprintf("There was an error in initializing the database. Error: %s ",err.Error())
		log.Printf("LOG_ERROR::INIT::InitializeDataBase::MSG:: %s", errorString)
		return errors.New(errorString)
	}
	return nil
}

func settingEnv() error{
	listSSM := []string{"JWT_SECRET"}
	// SSM parameter stores
	client, err := aws.NewSSMConfig()
	if err != nil {
		log.Printf("%s", "LOG_ERROR::INIT::SettingEnv::MSG::There was an error in creating the connection with AWS.")
	}

	for _, v := range listSSM {
		log.Printf("LOG_INFO::INIT::SettingEnv::MSG::Getting parameter with value: %s", v)
		ssmValue, err := client.GetParameterSSM(v)
		if err != nil {
			errorString:= fmt.Sprintf("Error getting the parameter store. Error: %s ",err.Error())
			log.Printf("LOG_ERROR::INIT::SettingEnv::MSG:: %s", errorString)
			return errors.New(errorString)
		}
		log.Printf("LOG_INFO::INIT::MSG::Setting enviroment variable with value: %s:%s", v, ssmValue)
		if err = os.Setenv(v, ssmValue); err != nil {
			errorString:= fmt.Sprintf("Error setting the enviroment variables. Error: %s ",err.Error())
			log.Printf("LOG_ERROR::INIT::SettingEnv::MSG:: %s", errorString)
			return errors.New(errorString)
		}
	}
	return nil
}
