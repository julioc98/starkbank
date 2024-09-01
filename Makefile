.PHONY: migrate-up
migrate-up:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up
