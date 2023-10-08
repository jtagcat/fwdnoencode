FROM golang:1.21 AS build

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/fwdnoencode

FROM gcr.io/distroless/static-debian11
LABEL org.opencontainers.image.source="https://github.com/jtagcat/calrefwdnoencodedact"

COPY --from=build /go/bin/fwdnoencode /

CMD ["/fwdnoencode"]
EXPOSE 8080
