PROTO_SRC = ./pkg/api/proto/matching_engine.proto
PROTO_DST = ./

.PHONY: proto
proto:
	protoc --go_out=$(PROTO_DST) --go-grpc_out=$(PROTO_DST) --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(PROTO_SRC)

go-test:
	go test ./test/... -v