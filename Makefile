.PHONY: run, test

run:
	docker-compose up

test:
	docker run --name=redis-test -p 6380:6379 -d --rm redis
	docker run --name=postgres-test -e POSTGRES_USER='postgres' -e POSTGRES_PASSWORD='postgres' -e POSTGRES_DB='postgres-test' -p 5438:5432 -d --rm postgres
	go test -count=1 ./tests/
	docker stop redis-test
	docker stop postgres-test
