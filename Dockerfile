FROM golang:alpine3.20 as builder

WORKDIR /go/src/app
COPY . .

RUN go mod download && CGO_ENABLED=0 go build /go/src/app

FROM gcr.io/distroless/static-debian12

COPY --from=builder /go/src/app /
CMD ["/socks5-server"]
