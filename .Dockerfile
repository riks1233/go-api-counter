# syntax=docker/dockerfile:1

# Building part.
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
# go mod download is not strictly needed, but is a good practice to have. But go build can fetch required dependencies.
RUN go mod download
# CGO_ENABLED=0 GOOS=linux disallows for go to depend on external C libraries.
# This means that the binary will be self-contained AFAIK.
# Without this, I had an error running the container because of some lib was missing.
RUN CGO_ENABLED=0 GOOS=linux go build -o appbuild .

# Running part (use `scratch` image as a lightweight container to run the Go binary.)
# If built within the same golang image container, the resulting image weighs 1.3GB.
# With `scratch` the image is just ~12MB. :mindblown:
FROM scratch
COPY --from=builder /app/appbuild /appbuild
# Turns out this is strictly for documentationn purposes. -p 8080:8080 would work without this anyways.
EXPOSE 8080
CMD ["/appbuild"]
