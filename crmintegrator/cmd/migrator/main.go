package main

import (
	"flag"
	"log"
	"crmintegrator/internal/config"

	// just to make things easier for deployment, not crucial to the logic itself, hope that's ok :)
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	direction := flag.String("migrate", "up", "specify if it should be `up` or `down` migration")
	flag.Parse()

	if *direction != "down" && *direction != "up" {
		log.Fatal("-migrate accepts [up, down] values only")
	}

	m, err := migrate.New("file://migrations", cfg.DBConnString)
	if err != nil {
		log.Fatal(err)
	}

	if *direction == "up" {
		if err := m.Up(); err != nil {
			if err != migrate.ErrNoChange {
				log.Fatal(err)
			}
		}
	}

	if *direction == "down" {
		if err := m.Down(); err != nil {
			if err != migrate.ErrNoChange {
				log.Fatal(err)
			}
		}
	}
}
