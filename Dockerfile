# Use an official Go runtime as a parent image
FROM golang:1.16-alpine3.14

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the Go application when the container starts
CMD ["/app/main"]
