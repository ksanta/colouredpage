# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Karl Santa"

# Set the Current Working Directory inside the container
WORKDIR /go/src/app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go get -d -v ./...
RUN go install -v ./...

# Expose port 8080 to the outside world
EXPOSE 8080

# Default command if none is given. Override to choose different colours.
CMD ["app", "green"]
