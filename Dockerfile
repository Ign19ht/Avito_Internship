FROM golang:1.19 as build-backend

WORKDIR /build
COPY . .
RUN go build -o app ./backend/main.go

FROM ubuntu:latest as prod

WORKDIR /prod

COPY --from=build-backend /build/app ./app

CMD ["/prod/app"]