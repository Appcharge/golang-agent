FROM golang:alpine AS builder

ENV GO111MODULE=on

WORKDIR $GOPATH/app/

COPY . .

RUN go mod download

EXPOSE ${APP_PORT}

CMD ["go", "run", "cmd/main.go"]