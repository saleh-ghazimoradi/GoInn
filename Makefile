ifneq (,$(wildcard ./app.env))
	include app.env
	export $(shell sed 's/=.*//' app.env)
endif


format:
	@echo "Applying go fmt to the project"
	go fmt ./...


vet:
	@echo "Checking for errors with vet"
	go vet ./...

dockerup:
	docker compose --env-file app.env up -d

dockerdown:
	docker compose --env-file app.env down

# Run the HTTP server
http:
	go run . http

# Declare targets that are not files
.PHONY: format vet dockerup dockerdown http