username := postgres
password := postgres
host := localhost
port := 5432
dbname := test
docker-container := psql-dev
BINARY_NAME := oscar

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

build:
	@echo "Building back end..."
	go build -o ${BINARY_NAME} main.go
	@echo "Binary built!"

run: build
	@echo "Starting back end..."
	@env ./${BINARY_NAME} &
	@echo "Back end started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

start: run

stop:
	@echo "Stopping back end..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped back end!"

restart: stop start

.PHONY: createdb dropdb migrateup migratedown sqlc run clean start stop restart