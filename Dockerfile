############################
# STEP 1 build executable binary
############################
FROM golang:alpine as builder
# Get source
WORKDIR $GOPATH/src/
COPY . .
# Fetch dependencies.
RUN apk update && apk add --no-cache git
RUN cd ./cmd/api/ && go get -d -v
# Build the binary.
RUN cd ./cmd/api/ && CGO_ENABLED=0 go build -a -ldflags '-s' -o /go/bin/pantori
# Create non-root user
RUN adduser pantori --disabled-password
############################
# STEP 2 get certificates
############################
FROM alpine:latest as certs
RUN apk --update add ca-certificates
############################
# STEP 3 build a small image
############################
FROM scratch
# Copy needed files.
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/bin/pantori /go/bin/pantori
COPY config.json /go/bin/config.json
COPY --from=builder /etc/passwd /etc/passwd
# Run.
USER pantori
CMD ["/go/bin/pantori"]


