FROM golang:1.23.0-alpine3.20 as builder
WORKDIR /app
COPY . /app
RUN go build -o /app/main /app/main.go

FROM alpine:3.20.2
WORKDIR /app
COPY --from=builder ./app .
CMD /app/main