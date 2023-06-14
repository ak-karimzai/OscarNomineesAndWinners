username := postgres
password := postgres
host := localhost
port := 5432
dbname := test
docker-container := psql-dev

createdb:
	sudo docker exec -it ${docker-container} createdb --username=${username}  --owner=${username} ${dbname}

dropdb:
	sudo docker exec -it ${docker-container} dropdb --username=${username} --if-exists ${dbname}

container:
	sudo docker run --name ${docker-container} -d -p 5432:5432 -e POSTGRES_PASSWORD=${password} -v pgdata:/var/lib/postgresql/data -d postgres:14.7-alpine

migrateup:
	cd dbScripts/migrations && goose postgres  postgres://${username}:${password}@${host}:${port}/${dbname} up

migratedown:
	cd dbScripts/migrations && goose postgres postgres://${username}:${password}@${host}:${port}/${dbname} down

sqlc:
	sqlc generate

.PHONY: createdb dropdb migrateup migratedown sqlc