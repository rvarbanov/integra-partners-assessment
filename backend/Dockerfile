FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

FROM alpine:latest AS runtime

WORKDIR /app

COPY --from=builder /app/scripts/wait-for.sh .
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./wait-for.sh", "db:5432", "--", "./main"]
