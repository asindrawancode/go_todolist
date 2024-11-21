# Use the official Golang image as the base
FROM golang:1.23.2 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire code from the host to the container
COPY . .

# Build the Go application
RUN go build -o todolist main.go

# Use a minimal base image for the final build
FROM gcr.io/distroless/base

# Copy the binary from builder
COPY --from=builder /app/todolist /todolist

# Expose the port the app runs on
EXPOSE 8000

# Command to run the executable
CMD ["/todolist"]
