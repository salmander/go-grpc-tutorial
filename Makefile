deps: generate
	go mod vendor

generate:
	go generate -v ./...

run_server:
	go run rpc_server/main.go

run_client:
	go run rpc_client/main.go
