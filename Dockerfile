FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build minimal static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notification ./cmd/server

# ---------------- PRODUCTION IMAGE ----------------
FROM gcr.io/distroless/base-debian11 AS final

WORKDIR /app

COPY --from=build /app/notification /app/notification

# Expose WS + HTTP port
EXPOSE 8081

ENTRYPOINT ["/app/notification"]
