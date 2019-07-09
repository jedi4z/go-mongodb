FROM golang:stretch AS build

# Set the directory where we will copy and build
WORKDIR /go/src/github.com/jedi4z/go-mongodb

# Copy the files needed into the container
COPY . .

# Install dependencies
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-mongodb .


# Use the alpine image for running the service
FROM alpine:latest

# Download dependencies and set the directory
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the application binary from the build stage into the container
COPY --from=build /go/src/github.com/github.com/jedi4z/go-mongodb .

# Run the service application
CMD ["./go-mongodb"]
