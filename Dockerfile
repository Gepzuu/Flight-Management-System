# Use the official Go image as a builder stage
FROM golang:1.18 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module file
COPY go.mod ./

# If your project eventually uses external dependencies, uncomment the next line
# RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

# Use a scratch image for the final stage
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/myapp /myapp

# Specify the command to run your application
CMD ["/myapp"]