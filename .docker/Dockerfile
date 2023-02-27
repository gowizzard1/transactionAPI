# Start with a base Golang image
FROM golang:1.17-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -o main .

# Expose port 8080 for the app to listen on
EXPOSE 8080

# Run the app when the container starts
CMD ["./main"]
