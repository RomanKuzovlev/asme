PROTO_SRC = ./pkg/api/proto/matching_engine.proto
PROTO_DST = ./

DOCKER_IMAGE_NAME=matching-engine
DOCKER_IMAGE_TAG=latest
DOCKERFILE_PATH=./
CONTAINER_PORT=50051

.PHONY: proto
proto:
	protoc --go_out=$(PROTO_DST) --go-grpc_out=$(PROTO_DST) --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(PROTO_SRC)

go-test:
	go test ./test/... -v

run:
	go run ./cmd/.

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) $(DOCKERFILE_PATH)

docker-run:
	docker run -p $(CONTAINER_PORT):$(CONTAINER_PORT) --env-file=$(DOCKERFILE_PATH)/.env $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

docker-clean:
	-docker stop $(DOCKER_IMAGE_NAME)
	-docker rm $(DOCKER_IMAGE_NAME)

docker-restart: docker-clean docker-build docker-run

update:
	eval $$(minikube docker-env) && \
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) $(DOCKERFILE_PATH) && \
	kubectl rollout restart deployment/$(DOCKER_IMAGE_NAME)