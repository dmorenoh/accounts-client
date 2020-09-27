FROM golang:1.12-alpine

RUN apk update && apk upgrade && \  
    apk add --no-cache bash git openssh && \
    apk add --no-cache util-linux

ENTRYPOINT export UUID=`uuidgen` 

