FROM golang:1.12-alpine

# Install git
RUN apk update && apk upgrade && \  
    apk add --no-cache bash git openssh

ENTRYPOINT export UUID=`uuidgen` && echo $UUIDFROM alpine:1.12

# Set working directory
WORKDIR /go/accounts-client

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o main .
# RUN go test ./...
# Run tests
CMD CGO_ENABLED=0 go test ./...