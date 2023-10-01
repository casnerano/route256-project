FROM golang:1.21

WORKDIR /app

ADD ./bin/app ./cart
ADD ./configs ./configs
ADD ./web ./web
ADD ./api/v1/openapiv2 ./api/v1/openapiv2

CMD ["/app/cart"]
