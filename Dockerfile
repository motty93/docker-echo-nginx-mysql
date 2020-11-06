ARG GO_VERSION=1.14.6
ARG ALPINE_VERSION=3.12

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as build
ENV APP_ROOT /go/src/app
ENV GO111MODULE=on

WORKDIR ${APP_ROOT}
COPY ./src .

RUN apk update \
    && apk add --no-cache git gcc libc-dev \
    && git config --global http.postBuffer 524288000 \
    && rm -rf /var/cache/apk/*

# package install
RUN go build -o tmp/main cmd/main.go \
    && go get -u github.com/cosmtrek/air \
    && go get -v github.com/rubenv/sql-migrate/... \
    && go get -u gorm.io/gorm \
    && go get -u gorm.io/driver/mysql

EXPOSE 8020
CMD ["air", "cmd/main.go"]
