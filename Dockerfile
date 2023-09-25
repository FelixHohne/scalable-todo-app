FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download and cache the Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o myapp cmd/main/main.go

# Expose a port that the application will listen on (modify as needed)
EXPOSE 8080

# Define the command to run your application
CMD ["./myapp"]
