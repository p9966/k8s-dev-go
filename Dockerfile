FROM golang:1.17.6

RUN apk add --no-cache git

WORKDIR /go/src/app

COPY . .

RUN go mod tidy
RUN go get -d -v ./...
RUN go build -o /go/bin/k8s-dev-go

EXPOSE 80

CMD ["/go/bin/k8s-dev-go"]
