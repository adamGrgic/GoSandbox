# Use an official Golang image as the base image
FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy the Go module file
COPY go.mod ./

# Download dependencies (if there are any)
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o /go-container

# Command to run the application
CMD ["/go-container"]
