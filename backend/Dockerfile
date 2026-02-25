ARG GO_VERSION=1.24
ARG APP_NAME=TaskManagerBackend

### Builder image
FROM golang:${GO_VERSION}-alpine AS builder
WORKDIR /src

# Install tools needed for fetching modules
RUN apk add --no-cache git ca-certificates

# Cache modules
COPY go.mod go.sum ./
RUN go mod download

# Copy sources and build
COPY . .
# Build a static binary for Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /${APP_NAME} .

### Final minimal image
FROM alpine:3.18 AS runtime
# Create an unprivileged user
RUN addgroup -S app && adduser -S -G app app && apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /${APP_NAME} /usr/local/bin/${APP_NAME}
RUN chown app:app /usr/local/bin/${APP_NAME} && chmod +x /usr/local/bin/${APP_NAME}

EXPOSE 8080
ENV PORT=8080
USER app
ENTRYPOINT ["/usr/local/bin/app"]