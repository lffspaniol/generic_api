FROM golang

COPY . $GOPATH/src/github.com/lffspaniol/generic_api


WORKDIR $GOPATH/src/github.com/lffspaniol/generic_api
COPY go.mod go.sum $GOPATH/src/github.com/lffspaniol/generic_api/
RUN go mod download
RUN go install 

EXPOSE 50051

CMD ["generic_api"]
