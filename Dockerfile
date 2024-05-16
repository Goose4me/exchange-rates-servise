# Use the official Golang image as the base image
FROM golang:1.18-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o gses2.app .

# Start a new stage from scratch
FROM alpine:latest  

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/gses2.app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./gses2.app"]
