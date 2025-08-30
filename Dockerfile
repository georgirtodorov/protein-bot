# syntax=docker/dockerfile:1

FROM golang:1.25.0-alpine AS build

# Set working directory to module root
WORKDIR /app/cmd

# Copy module files and download dependencies
COPY cmd/go.mod cmd/go.sum ./
RUN go mod download

# Copy the main.go file
COPY cmd/main.go .

# Build the binary directly from main.go
RUN go build -o /app/protein-bot main.go

# ---- Final image ----
FROM scratch

COPY --from=build /app/protein-bot /protein-bot

CMD ["/protein-bot"]
