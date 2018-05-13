# minimal linux distribution
FROM golang:1.10-alpine

# GO and PATH env variables already set in golang image
# to reduce download time
RUN apk add -U make git

# set the go path to import the source project
WORKDIR $GOPATH/src/github.com/Sfeir/handsongo
ADD . $GOPATH/src/github.com/Sfeir/handsongo

# In one command-line (for reduce memory usage purposes),
# we install the required software,
# we build handsongo program
# we clean the system from all build dependencies
RUN make all && apk del make git && \
  rm -rf /go/pkg && \
  rm -rf /go/src && \
  rm -rf /var/cache/apk/* &&\
  mkdir /app && mv /go/bin/handsongo /app

# by default, the exposed ports are 8020 (HTTP)
EXPOSE 8020
ENTRYPOINT ["/app/handsongo"]
