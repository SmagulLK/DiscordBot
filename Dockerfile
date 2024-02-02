
FROM golang:alpine as builder


WORKDIR /app
COPY go.mod go.sum ./


RUN go mod download


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app cmd/main.go


FROM alpine:latest

WORKDIR /root/


COPY --from=builder /app/app .

RUN chmod +x ./app

EXPOSE 8080


CMD ["./app"]