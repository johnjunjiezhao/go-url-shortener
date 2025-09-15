# Multi-stage build for the Go URL Shortener

FROM golang:1.25.1 AS build
WORKDIR /app

# Cache deps first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build a static binary (multi-arch friendly)
RUN CGO_ENABLED=0 GOOS=linux go build -o /server .

# Minimal runtime image
FROM scratch AS final
WORKDIR /
COPY --from=build /server /server

EXPOSE 9808
ENTRYPOINT ["/server"]
