# # Makefile для создания миграций

# # Переменные которые будут использоваться в наших командах (Таргетах)
# DB_DSN := "postgres://max:@localhost:5432/main?sslmode=disable"
# MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# # Таргет для создания новой миграции
# migrate-new:
# 	migrate create -ext sql -dir ./migrations ${NAME}

# # Применение миграций
# migrate:
# 	$(MIGRATE) up

# # Откат миграций
# migrate-down:
# 	$(MIGRATE) down
	
# # для удобства добавим команду run, которая будет запускать наше приложение
# run:
# 	go run cmd/app/main.go 

	





# Переменные, которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://max:@localhost:5432/main?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Указываем, что таргеты не являются файлами
.PHONY: migrate-new migrate migrate-down run

# Таргет для создания новой миграции
# Использование: make migrate-new NAME=<имя_миграции>
migrate-new:
	migrate create -ext sql -dir ./migrations -seq $(NAME)

# Применение миграций
migrate:
	@echo "Applying migrations..."
	$(MIGRATE) up

# Откат миграций
migrate-down:
	@echo "Reverting migrations..."
	$(MIGRATE) down

# Таргет для запуска приложения
run:
	@echo "Starting the application..."
	go run cmd/app/main.go