.PHONY: protoc
protoc:
	protoc --go_out=. --go-grpc_out=. proto/pingpong.proto


.PHONY: build_ping
build_ping:
	go build -o ./bin/ping_service ./cmd/ping_service

.PHONY: build_pong
build_pong:
	go build -o ./bin/pong_service ./cmd/pong_service

.PHONY: build_reflector
build_reflector:
	go build -o ./bin/reflector ./cmd/reflector

.PHONY: build
build: build_ping build_pong build_reflector

.PHONY: docker_build
docker_build: docker_build_ping docker_build_pong docker_build_reflector

.PHONY: docker_build_ping
docker_build_ping: 
	docker build --target ping_service -t ping:latest .

.PHONY: docker_build_pong
docker_build_pong: 
	docker build --target pong_service -t pong:latest .

.PHONY: docker_build_reflector
docker_build_reflector: 
	docker build --target reflector_service -t reflector:latest .
