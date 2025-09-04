postgres:
	docker run --name postgres13 -p 5433:5432 -e POSTGRES_PASSWORD=Franzego@1 -e POSTGRES_USER franz -d postgres:13.22-alpine3.22

createdb:
	docker exec -it postgres13 createdb -U franz ecommerce

dropdb:
	docker exec -it postgres13 dropdb -U franz ecommerce

migrateup:
	docker exec -it postgres13 migrate -path /db/migration -database "postgresql://franz:Franzego@1@localhost:5433/new?sslmode=disable" -verbose up

migratedown:
	docker exec -it postgres13 migrate -path /db/migration -database "postgresql://franz:Franzego@1@localhost:5433/new?sslmode=disable" -verbose down

sqlc:
	sqlc generate


.PHONY: postgres createdb dropdb migrateup migratedown sqlc
