FROM golang:alpine as builder

# Install dependencies
RUN apk add --update git build-base
RUN go get -u github.com/golang/dep/cmd/dep

# Copy source
COPY main.go Gopkg.toml ${GOPATH}/src/github.com/dthongvl/cinerum/
COPY src/ ${GOPATH}/src/github.com/dthongvl/cinerum/src/
WORKDIR ${GOPATH}/src/github.com/dthongvl/cinerum/

# Build
RUN dep ensure
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o main

FROM alpine:latest

RUN apk add --update ffmpeg

WORKDIR /root/
COPY --from=builder /go/src/github.com/dthongvl/cinerum/main .
COPY static ./static
COPY template ./template
RUN mkdir -p preview

EXPOSE 3000
ENTRYPOINT ["./main"]