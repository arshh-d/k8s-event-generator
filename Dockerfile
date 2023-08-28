# Use an official Go runtime as a parent image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the Go application
RUN go build -o event-generator .

# Set environment variables if needed
# ENV VAR_NAME=value

# Run the application
CMD ["./event-generator"]
