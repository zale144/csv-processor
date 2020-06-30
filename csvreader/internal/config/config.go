package config

import (
	"errors"
	"log"
	"os"
	"strconv"
)

type Config struct {
	HTTPPort,
	BatchSize int64
	GRPCURL string
}

func Get() (*Config, error) {

	grpcHost := os.Getenv("GRPC_HOST")
	grpcPort := os.Getenv("GRPC_PORT")
	httpPortStr := os.Getenv("CSV_HTTP_PORT")
	batchSizeStr := os.Getenv("BATCH_SIZE")

	var (
		batchSize int64 = 100
		httpPort  int64
	)

	if grpcHost == "" {
		return nil, errors.New("env GRPC_HOST missing")
	}
	if grpcPort == "" {
		return nil, errors.New("env GRPC_PORT missing")
	}

	if httpPortStr == "" {
		return nil, errors.New("env CSV_HTTP_PORT missing")
	}

	hp, err := strconv.ParseInt(httpPortStr, 10, 32)
	if err != nil {
		log.Println(err)
	} else {
		httpPort = hp
	}

	if batchSizeStr != "" {
		bs, err := strconv.ParseInt(batchSizeStr, 10, 32)
		if err != nil {
			log.Println(err)
		} else {
			batchSize = bs
		}
	}

	return &Config{
		HTTPPort:  httpPort,
		BatchSize: batchSize,
		GRPCURL:   grpcHost + ":" + grpcPort,
	}, nil
}
