# stage development
FROM golang:1.14 as development

WORKDIR /go/src/github.com/mniudanri/store
COPY . .

RUN go get golang.org/x/lint/golint
RUN go get github.com/swaggo/swag/cmd/swag
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go","run","app"]
