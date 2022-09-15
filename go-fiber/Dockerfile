# syntax=docker/dockerfile:1
FROM golang:1.19.1-alpine3.16 as build

ENV CGO_ENABLED=0

WORKDIR /app

RUN go mod init main

RUN go get github.com/gofiber/fiber/v2

COPY . .

RUN go build -a -ldflags '-extldflags "-static"'

FROM gcr.io/distroless/static-debian11

COPY --from=build /app/main ./

CMD ["./main"]

EXPOSE 3000