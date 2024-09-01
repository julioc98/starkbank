.PHONY: migrate-up
migrate-up:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up


.PHONY: remote-migrate-up
remote-migrate-up:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database '34.151.242.116://postgres:postgres@localhost:5432/postgres?sslmode=disable' up

