build:
	docker build . -t tender-service

run: build
	docker run -p 8080:8080 tender-service

init-db:
	@echo "==> Wait for the database to be initialized..."
	docker compose exec db psql -U postgres -d tender-service -f /docker-entrypoint-initdb.d/create_table.sql

docker-build:
	@echo "==> Docker containers are being built..."
	docker compose build

docker-up: docker-build
	@echo "==> Docker containers are starting..."
	docker compose up

docker-down:
	@echo "==> Docker containers are stopping their work..."
	docker compose down