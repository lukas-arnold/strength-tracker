FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 go build -o app .


FROM alpine:3.20

WORKDIR /root/
COPY --from=builder /app/app .

CMD ["./app"]