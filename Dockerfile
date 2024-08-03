# Use the official Golang image based on Alpine Linux
FROM golang:alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Install git and other dependencies
RUN apk update && apk upgrade && \
    apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Expose port 8080 to the outside world
EXPOSE 8001

# Command to run the executable
CMD ["go", "run", "main.go"]
