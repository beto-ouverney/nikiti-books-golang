FROM golang

WORKDIR /app

RUN apt-get update
RUN apt-get install lsof

COPY . .

RUN go get -d -v ./...

EXPOSE 8080