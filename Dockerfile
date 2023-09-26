# Use the official Go image to create a build artifact.
FROM golang:1.18 AS builder

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Copy local code to the container image.
COPY . ./

# Build the command inside the container.
RUN go build -o server ./cmd/server

# Use a Docker multi-stage build to create a lean production image.
FROM gcr.io/distroless/base-debian10

# Copy the binary to the production image from the builder stage.
COPY --from=builder /src/server /server

# Run the web service on container startup.
CMD ["/server"]
