FROM golang:1.13

WORKDIR /go/src/
COPY gomodule /go/src/gomodule

WORKDIR /go/src/gomodule

# Needed because you cannot copy source files under /go/src
ENV GO111MODULE=on

# This is not needed if you are not behind a firewall
ENV http_proxy=http://proxy.houston.hpecorp.net:8080
ENV https_proxy=http://proxy.houston.hpecorp.net:8080

RUN go mod tidy
RUN go mod download
RUN go mod verify
CMD ["go", "run", "main.go"]
