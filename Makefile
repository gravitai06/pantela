DB_DSN := "postgres://max:@localhost:5432/main?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

generate-users:
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/users.yaml > ./internal/web/users/api.gen.go

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/tasks.yaml > ./internal/web/tasks/api.gen.go

lint:
	golangci-lint run --out-format=colored-line-number

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down
	
run:
	go run cmd/app/main.go 
