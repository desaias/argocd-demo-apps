ARG GOLANG_VERSION=1.22.1
FROM golang:${GOLANG_VERSION} AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build each app
RUN CGO_ENABLED=0 GOOS=linux go build -o hello ./cmd/hello && \
CGO_ENABLED=0 GOOS=linux go build -o goodbye ./cmd/goodbye

FROM alpine:latest
WORKDIR /apps
COPY --from=builder /app/hello .
COPY --from=builder /app/goodbye .
