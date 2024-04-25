FROM golang:1.22-alpine as build


WORKDIR /app

COPY go.mod go.sum ./


RUN go mod download

COPY . .


RUN GOOS=linux go build .

FROM alpine:latest as runtime

WORKDIR /app


COPY --from=build /app/sample-api /app/sample-api

CMD ["/app/sample-api"]