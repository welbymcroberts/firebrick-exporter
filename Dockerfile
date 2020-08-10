FROM golang:alpine as builder
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
ENV USER=webapp
ENV UID=1001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    webapp
WORKDIR $GOPATH/src/welbymcroberts/firebrick-exporter/
COPY . .
RUN go mod download
RUN go mod verify
RUN go build -ldflags="-w -s" -o /go/bin/firebrick-exporter
RUN chmod +x /go/bin/firebrick-exporter

#############################
FROM scratch
# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /lib/ld-musl-x86_64.so.1 /lib/ld-musl-x86_64.so.1
COPY --from=builder /go/bin/firebrick-exporter /go/bin/firebrick-exporter
# TODO - config.yaml
USER webapp:webapp
ENTRYPOINT ["/go/bin/firebrick-exporter"]