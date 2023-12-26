FROM docker.io/library/golang:1.21.5-bullseye

RUN go install golang.org/x/tools/gopls@latest &&\ 
    go install github.com/go-delve/delve/cmd/dlv@latest &&\
    go install honnef.co/go/tools/cmd/staticcheck@latest