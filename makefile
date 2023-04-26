docker-build:
	docker-compose build

docker-run:
	docker-compose up

docker-down:
	docker-compose down

docker-clear-container:
	docker rm $(docker container ls -a -q)