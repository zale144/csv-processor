package main

import (
	"csvreader/internal/api/server"
	"csvreader/internal/app/csv"
	"csvreader/internal/config"

	"log"
)

func main() {

	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	grpcConn, err := server.ServeGrpc(cfg.GRPCURL)
	if err != nil {
		log.Fatal(err)
	}

	up := csv.NewUserProcessor(grpcConn, cfg.BatchSize)
	server.ServeHTTP(cfg.HTTPPort, up)
}
