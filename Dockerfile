FROM golang:1.16-alpine as build

RUN apk add --no-cache git

RUN go get -v github.com/cosmtrek/air

WORKDIR /src

COPY . /src

RUN go mod vendor

COPY docker/app/entrypoint.sh /usr/local/bin/

RUN chmod +x /usr/local/bin/entrypoint.sh

EXPOSE ${SITE_PORT}

CMD ["entrypoint.sh"]