FROM golang:alpine AS builder



# Move to working directory /src
WORKDIR /src


# Copy the code into the container
COPY wildcardproxy.go /src


RUN CGO_ENABLED=0 go build -o /bin/wildcardproxy


# Build a small image
FROM scratch

COPY --from=builder /bin/wildcardproxy /

# Command to run
ENTRYPOINT ["/wildcardproxy"]