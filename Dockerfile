############################
# STEP 1 build executable binary
############################
FROM golang:alpine as builder
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/
COPY . .
# Fetch dependencies.
# Using go get.
RUN cd ./cmd/api/ && go get -d -v
# Build the binary.
RUN cd ./cmd/api/ && CGO_ENABLED=0 go build -a -ldflags '-s' -o /go/bin/pantori
# Create user
RUN adduser pantori --disabled-password
############################
# STEP 2 include
############################
FROM alpine:latest as certs
RUN apk --update add ca-certificates
############################
# STEP 3 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/bin/pantori /go/bin/pantori
COPY config.json /go/bin/config.json
COPY --from=builder /etc/passwd /etc/passwd
# Run the hello binary.
USER pantori
CMD ["/go/bin/pantori"]