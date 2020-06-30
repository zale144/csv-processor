package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	GRPCURL,
	DBConnString,
	CRMURL string
	RetryLimit uint
	RetryDelay time.Duration
}

func Get() (*Config, error) {

	dbPort := os.Getenv("POSTGRES_PORT")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	gRPCHost := os.Getenv("GRPC_HOST")
	gRPCPort := os.Getenv("GRPC_PORT")
	retryLimit := os.Getenv("RETRY_LIMIT")
	retryDelay := os.Getenv("RETRY_DELAY")

	cfg := &Config{
		CRMURL:   os.Getenv("CRM_URL"),
		RetryLimit: 10,
		RetryDelay: 100 * time.Millisecond,
	}

	if retryLimit != "" {
		if retLimit, err := strconv.ParseInt(retryLimit, 10, 32); err != nil {
			log.Println(err)
		} else {
			cfg.RetryLimit = uint(retLimit)
		}
	}

	if retryDelay != "" {
		if retDelay, err := strconv.ParseInt(retryDelay, 10, 32); err != nil {
			log.Println(err)
		} else {
			cfg.RetryDelay = time.Duration(retDelay) * time.Millisecond
		}
	}

	if gRPCHost == "" {
		return nil, errors.New("env GRPC_HOST missing")
	}
	if gRPCPort == "" {
		return nil, errors.New("env GRPC_PORT missing")
	}
	if dbHost == "" {
		return nil, errors.New("env POSTGRES_HOST missing")
	}
	if dbPort == "" {
		return nil, errors.New("env POSTGRES_PORT missing")
	}
	if dbUser == "" {
		return nil, errors.New("env POSTGRES_USER missing")
	}
	if dbname == "" {
		return nil, errors.New("env POSTGRES_DB missing")
	}
	if cfg.CRMURL == "" {
		return nil, errors.New("env CRM_URL missing")
	}


	cfg.DBConnString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbname, "disable" )

	cfg.GRPCURL = gRPCHost + ":" + gRPCPort

	return cfg, nil
}
