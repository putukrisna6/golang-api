FROM golang:1.17.8-alpine

LABEL maintainer="Putu Krisna Andyartha"

RUN apk update && apk add --no-cache git

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /bin/myapp

CMD [ "/bin/myapp" ]