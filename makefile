CONTAINER_NAME=go-api

docker-build:
	docker-compose build

docker-run:
	docker-compose up

docker-down:
	docker-compose down

docker-clear-container:
	docker rm $(docker container ls -a -q)

run-tests:
	docker exec -t ${CONTAINER_NAME} go test ./...

run-tests-with-coverage:
	docker exec -t ${CONTAINER_NAME} go test ./... --coverprofile coverage.txt