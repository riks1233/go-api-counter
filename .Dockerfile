# syntax=docker/dockerfile:1
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o appbuild .

FROM scratch
COPY --from=builder /app/appbuild /appbuild
EXPOSE 8080
CMD ["/appbuild"]
