.PHONY: first-start
#only first start project
init-dev-db:
	createdb reastapi_dev
init-test-db:
	createdb reastapi_test

.PHONY: migrate 
migrate-u:
	migrate -database "postgresql://localhost:5432/rtlabs?sslmode=disable" -path migrations up
	migrate -database "postgresql://localhost:5432/rtlabs_test?sslmode=disable" -path migrations up

migrate-d:
	migrate -database "postgresql://localhost:5432/rtlabs?sslmode=disable" -path migrations down
	migrate -database "postgresql://localhost:5432/rtlabs_test?sslmode=disable" -path migrations down

migrate-force:
	migrate -path migrations -database "postgresql://localhost:5432/rtlabs?sslmode=disable" force __VERCION__

.PHONY: build
build:
	go build -v ./cmd/apiserver

run:
	go build -v ./cmd/apiserver
	./apiserver

.PHONY: test
test:
	go test -race -cover -timeout 30s ./internal/app/...

test-v:
	go test -v -race -cover -timeout 30s ./internal/app/...

test-p:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

test-ph:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

.DEFAULT_GOAL := build

# need postgress@13
fix-bd:
	rm /usr/local/var/postgres/postmaster.pid
	brew services restart postgres
	postgres -v
