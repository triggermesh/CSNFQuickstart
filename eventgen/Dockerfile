FROM golang:1.17-buster AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /go/src/eventgen

COPY go.mod ./

COPY . .
RUN go build .

FROM registry.access.redhat.com/ubi8/ubi-minimal

COPY --from=builder /go/src/eventgen/eventgen /

ENTRYPOINT ["/eventgen"]
