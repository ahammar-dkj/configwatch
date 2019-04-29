FROM golang:1.12

WORKDIR $GOPATH/src/github.com/ahammar-dkj/configwatch

COPY . .
RUN go get -d -v ./... && go install -v ./...
RUN mkdir /etc/configwatch

CMD ["configwatch", "-d", "/etc/configwatch"]

