FROM golang


COPY . $GOPATH/src/github.com/lffspaniol/generic_api


ENV GO111MODULE on

WORKDIR $GOPATH/src/github.com/lffspaniol/generic_api
COPY go.mod go.sum $GOPATH/src/github.com/lffspaniol/generic_api/
RUN go get -d -v ./...
RUN go build -v -i -o generic_api  cmd/server/main.go 

EXPOSE 50051
EXPOSE 8080

CMD ["./generic_api"]
