generate:
	protoc --openapiv2_out=openapiv2 --grpc-gateway_out=. --go_out=. --go-grpc_out=. --proto_path=proto proto/*.proto