FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /catalog-api ./cmd/catalog-api/

FROM alpine:latest

RUN adduser -D -g '' appuser

USER appuser

COPY --from=builder /catalog-api /usr/local/bin/catalog-api

EXPOSE 8080

CMD [ "/usr/local/bin/catalog-api" ]
