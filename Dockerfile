FROM golang:1.20-alpine
WORKDIR /go/src/github.com/mxload/buuurst_dev_go

RUN apk add --no-cache gcc \
    git \
    libc-dev \
    make

COPY go.mod .
COPY go.sum .

RUN go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install golang.org/x/tools/gopls@latest \
    && go install golang.org/x/tools/cmd/goimports@latest \
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2
