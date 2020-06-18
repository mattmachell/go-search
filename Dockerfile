FROM golang:1.14.3-alpine3.11

RUN apk --update --no-cache add \
    ca-certificates \
    && rm -rf /var/cache/apk/*
    
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]