build:
	docker build -t gorm-migrate:latest -f docker/mariadb/Dockerfile .

run:
	docker compose up -d