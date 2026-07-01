FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

# GONOSUMCHECK=* skips checksum DB verification in air-gapped or proxy-restricted environments
# GONOSUMDB=* skips the checksum database lookup
# GOFLAGS=-mod=mod allows go to update go.mod and go.sum automatically
RUN GONOSUMCHECK=* GONOSUMDB=* GOFLAGS=-mod=mod go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -mod=mod -o /app/server ./cmd/api

FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
