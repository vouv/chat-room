FROM golang:alpine

MAINTAINER monigo

ENV PROJECT_PATH app

WORKDIR $GOPATH

COPY . $GOPATH/src/$PROJECT_PATH

COPY ./web $GOPATH/web

RUN go build -o app $GOPATH/src/$PROJECT_PATH/cmd/app.go

ENV PORT 8080
ENV GIN_MODE release

EXPOSE 8080

ENTRYPOINT ["sh", "-c", "./app"]