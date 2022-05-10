FROM golang:1.18-alpine as builder

LABEL maintainer="Alexander Miranda <alexandermichaelmiranda@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8080:80

CMD ["./main"]

