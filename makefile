all: build run

build:
	go build app.go main.go user_model.go widget_model.go

run:
	./app