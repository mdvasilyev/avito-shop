FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . /app
COPY ./migrations/ /app/migrations/

RUN CGO_ENABLED=0 go build -o app cmd/app/main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/migrations /app/migrations/

ENTRYPOINT [ "./app" ]
