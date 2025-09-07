# syntax=docker/dockerfile:1

FROM golang:1.25.0-alpine AS build

# Set working directory to module root
WORKDIR /app

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the main.go file and other source file dependencies
COPY cmd ./cmd/main.go
COPY internal ./internal

# Build the binary directly from main.go
RUN go build -o /app/protein-bot ./cmd/main.go

# ---- Final image ----
FROM scratch

COPY --from=build /app/protein-bot /protein-bot

EXPOSE 8080

CMD ["/protein-bot"]
