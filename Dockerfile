# Build Stage
# First pull Golang image
FROM golang:1.17-alpine as build-env

# Set environment variable
ENV APP_NAME simple-restful-bookstore
ENV CMD_PATH server.go

# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

RUN apk add build-base

RUN go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

RUN apk --update-cache add sqlite \
    && rm -rf /var/cache/apk/* \
    && sqlite3 $GOPATH/src/$APP_NAME/bookstore.sqlite < $GOPATH/src/$APP_NAME/migrations/init.sql \
    && chmod a+rw $GOPATH/src/$APP_NAME/bookstore.sqlite

# Run Stage
FROM alpine:3.14

# Set environment variable
ENV APP_NAME simple-restful-bookstore

# Copy only required data into this image
COPY --from=build-env /$APP_NAME .

# Expose application port
EXPOSE 1323

# Start app
CMD ./$APP_NAME