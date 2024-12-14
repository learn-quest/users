# Step 1: Build stage
FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o app .

# Set up environment
ENV GO_ENV=production

# Expose the port the app runs on
EXPOSE 4000

# Command to run the application
CMD ["./app"]
