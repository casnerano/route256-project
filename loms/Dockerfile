FROM golang:1.21

ADD ./bin/app /app
ADD ./configs /configs

CMD ["/app", "-config", "/configs/config.yaml"]
