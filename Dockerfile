FROM golang:1.19-alpine AS builder
ARG GITHUB_SHA
ARG GITHUB_REF
RUN apk add --no-cache ca-certificates tzdata make
WORKDIR /src
COPY go.mod go.sum Makefile ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
RUN make build

FROM scratch
USER 1000
WORKDIR /app
COPY --from=builder /src/bin/server /app/server
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
CMD ["/app/server"]
