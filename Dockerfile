FROM golang:1.14-alpine

LABEL maintainer="Alexander Miranda <alexandermichaelmiranda@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main /app/cmd/server/main.go

EXPOSE 8080

CMD ["./main"]
