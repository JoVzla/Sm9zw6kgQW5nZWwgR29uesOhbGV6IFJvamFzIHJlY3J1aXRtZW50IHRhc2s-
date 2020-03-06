# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.14
ARG GO_VERSION=1.14

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder
ENV GO111MODULE=on
RUN mkdir /app
RUN apk add --no-cache ca-certificates git
WORKDIR /app
COPY go.mod .
COPY go.sum .
# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/app .
COPY ./config.yml .
RUN ["chmod", "+x", "app"]
EXPOSE 9010
CMD ["./app"]
