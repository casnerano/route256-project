FROM golang:1.21

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /app

COPY /cart/migrations/postgres ./migrations/cart
COPY /loms/migrations/postgres ./migrations/loms

COPY /infrastructure/migration/start.sh ./start.sh
RUN chmod +x /app/start.sh

CMD ["/app/start.sh"]
