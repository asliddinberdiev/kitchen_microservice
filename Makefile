tidy:
	@go mod tidy
	@go mod vendor

run-orders:
	@go run services/orders/*.go

run-kitchen:
	@go run services/kitchen/*.go

gproto:
	@protoc \
		--proto_path=protos "protos/orders.proto" \
		--go_out=services/common/genproto/orders --go_opt=paths=source_relative \
		--go-grpc_out=services/common/genproto/orders \
		--go-grpc_opt=paths=source_relative