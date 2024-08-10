build:
	swag init
	go build -o strive_backend -ldflags=-X=main.version=${VERSION} .

run:
	swag init
	go run .

update-dockerfile:
	docker-env