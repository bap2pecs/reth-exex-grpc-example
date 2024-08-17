.PHONY: start-grpc build-docker demo ls logs query-block proto-gen

proto-gen:
	@echo "Generating protobuf files..."
	@cd grpc/proto && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative MinimalExExHandler.proto

start-grpc:
	@echo "Starting gRPC server..."
	@cd grpc && go run server/grpcserver.go

build-docker:
	@echo "Building Docker image for the reth-exex..."
	@cd exex && docker build . -t reth:fm

demo-run:
	@echo "Running demo fmnet..."
	@kurtosis run github.com/ethpandaops/ethereum-package --args-file ./demo/fm_network_params.yaml --image-download always --enclave fm

demo-stop:
	@echo "Removing demo fmnet..."
	@kurtosis enclave rm fm -f

demo-restart: demo-stop demo-run

ls-enclaves:
	@kurtosis enclave ls

ls:
	@kurtosis enclave inspect fm

logs:
	@kurtosis service logs fm el-2-reth-teku

# usages: make query-block rpc=127.0.0.1:54643
query-block:
	@echo "Querying latest block number..."
	@curl -X POST -H 'Content-Type: application/json' -d '{"jsonrpc": "2.0", "method": "eth_blockNumber", "params": [], "id": 1}' $(rpc)