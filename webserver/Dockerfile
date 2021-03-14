FROM golang:1.16.0

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY . .

RUN go mod download && go build webserver.go

FROM alpine:latest

COPY --from=0 /app .

EXPOSE 1323

CMD ["./webserver"]
