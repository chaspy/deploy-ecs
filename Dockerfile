FROM golang:1.18.0 as builder

WORKDIR /go/src

#COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download

COPY . .

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/hello \
    -ldflags '-s -w'

FROM alpine:3.14.4 as runner

COPY --from=builder /go/bin/hello /app/hello

RUN adduser -D -S -H hello

USER hello

ENTRYPOINT ["/app/hello"]

