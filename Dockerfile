FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/gabrielucelli/widget/api

# set the api as the working directory
WORKDIR /go/src/github.com/gabrielucelli/widget/api

# install sqlite3
RUN apt-get -y update
RUN apt-get install -y sqlite3

# populate the database
RUN make populate_db

# Install our dependencies
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/schema
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/mattn/go-sqlite3

# build the application
RUN make build

# Expose default port (3000)
EXPOSE 3000

# run the application
CMD ["make", "run"]