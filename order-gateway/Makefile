ORDER_GATEWAY_IMAGE_NAME=order-gateway
ORDER_GATEWAY_IMAGE_TAG=latest
ORDER_GATEWAY_PATH=./
ORDER_GATEWAY_PORT=8080

run:
	go run ./cmd/.

test-go:
	go run ./test/test_order_gateway.go

docker-build:
	docker build -t $(ORDER_GATEWAY_IMAGE_NAME):$(ORDER_GATEWAY_IMAGE_TAG) $(ORDER_GATEWAY_PATH)

docker-clean:
	-docker stop $(ORDER_GATEWAY_IMAGE_NAME)
	-docker rm $(ORDER_GATEWAY_IMAGE_NAME)

docker-restart: order-gateway-docker-clean order-gateway-docker-build order-gateway-docker-run

update:
	eval $$(minikube docker-env) && \
	docker build -t $(ORDER_GATEWAY_IMAGE_NAME):$(ORDER_GATEWAY_IMAGE_TAG) $(ORDER_GATEWAY_PATH) && \
	kubectl rollout restart deployment/$(ORDER_GATEWAY_IMAGE_NAME)