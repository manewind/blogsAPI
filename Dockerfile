# Use the official Go image as the base image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies first
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o myapp

# Expose the port your application will run on (e.g., 8080)
EXPOSE 8080

# Command to run the Go application
CMD ["./myapp"]
