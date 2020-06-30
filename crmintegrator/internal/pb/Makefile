.PHONY: all

all:
	@protoc --proto_path=. --go_out=plugins=grpc:. user.proto
	@protoc --proto_path=. --go_out=plugins=grpc:../../../csvReader/internal/pb user.proto
	@echo "Protofiles rebuild done"
