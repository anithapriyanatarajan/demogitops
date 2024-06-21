# Use the official Golang image as the base image
FROM golang:1.17-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application will run on
EXPOSE 8084

# Command to run the executable
CMD ["./main"]