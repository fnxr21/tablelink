
push:
	sh push.sh
run:
	go run main.go
dc:
	docker compose up
dc-up:
	docker compose up -d

dc-down:
	docker compose down
proto:
    protoc --go_out=protobuf \
    --go-grpc_out=protobuf --go-grpc_opt=require_unimplemented_servers=false\
    --grpc-gateway_out=protobuf \
    protobuf/item.proto
