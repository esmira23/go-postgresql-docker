build:
	docker-compose build
	
run:
	docker-compose up -d

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root golang-postgresql-docker

dropdb:
	docker exec -it postgres15 dropdb golang-postgresql-docker

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/golang-postgresql-docker?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/golang-postgresql-docker?sslmode=disable" -verbose down

.PHONY: build run postgres createdb dropdb migrateup migratedown

