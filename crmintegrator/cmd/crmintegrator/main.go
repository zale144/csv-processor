package main

import (
	"log"
	"crmintegrator/internal/api/server"
	"crmintegrator/internal/app/user"
	"crmintegrator/internal/client"
	"crmintegrator/internal/config"
	db "crmintegrator/pkg/database"
)

func main() {

	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	dB, err := db.Init(cfg)
	if err != nil {
		log.Fatal(err)
	}

	crmClient := client.NewCRMClient(cfg)
	usrSrvc := user.NewUserService(cfg, dB, crmClient)

	if err := server.Start(cfg, usrSrvc); err != nil {
		log.Fatal(err)
	}
}
