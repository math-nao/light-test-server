FROM golang:1.12 AS build
LABEL maintainer "math.nao@outlook.com"

WORKDIR /go/src/app
COPY main.go .

RUN CGO_ENABLED=0 go build -o dist/app .

FROM alpine:latest

WORKDIR /root/

COPY --from=build /go/src/app/dist/app .

EXPOSE 8080

CMD ["./app"]
