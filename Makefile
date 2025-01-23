# Переменные
DB_DSN := "postgres://max:@localhost:5432/main?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)
OPENAPI_PATH := openapi/openapi.yaml

# Генерация кода для пользователей
generate-users:
	oapi-codegen -generate types,server,strict-server -package tasks -o internal/web/tasks/api.gen.go openapi/openapi.yaml

# Генерация кода для задач
generate-tasks:
	oapi-codegen -generate types,server,strict-server -package tasks -o internal/web/tasks/api.gen.go openapi/openapi.yaml

# Генерация всего кода
generate-all: generate-tasks generate-users

# Линтинг
lint:
	golangci-lint run --out-format=colored-line-number

# Создание новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# Запуск приложения
run:
	go run cmd/app/main.go

# Очистка сгенерированных файлов
clean:
	rm -f internal/web/tasks/api.gen.go internal/web/users/api.gen.go

#DB_DSN := "postgres://max:@localhost:5432/main?sslmode=disable"
#MIGRATE := migrate -path ./migrations -database $(DB_DSN)
#
#generate-users:
#	oapi-codegen -generate types,server -package users -o internal/web/users/api.gen.go openapi/openapi.yaml
#
#gen:
#	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
#
#lint:
#	golangci-lint run --out-format=colored-line-number
#
#migrate-new:
#	migrate create -ext sql -dir ./migrations ${NAME}
#
#migrate:
#	$(MIGRATE) up
#
#migrate-down:
#	$(MIGRATE) down
#
#run:
#	go run cmd/app/main.go
