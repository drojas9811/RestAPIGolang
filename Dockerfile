FROM golang:1.19.0

COPY ./ /go/src/

WORKDIR /go/src/cmd

RUN go build -o ./bin/RESTAPI

CMD ["./bin/RESTAPI"]