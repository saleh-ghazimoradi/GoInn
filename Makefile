docker-up:
	docker compose up -d
docker-down:
	docker compose down
fmt:
	go fmt ./...
vet:
	go vet ./...
http: fmt vet
	go run . http