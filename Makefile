format:
	go fmt ./...

vet:
	go vet ./...

dockerUp:
	docker compose up -d

dockerDown:
	docker compose down

http: format vet
	go run . http

