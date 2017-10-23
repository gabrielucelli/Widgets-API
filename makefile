all: build run

build:
	go build app.go main.go user_model.go widget_model.go

populate_db:
	sh create_users_table.sh 
	sh create_widget_table.sh

run:
	./app