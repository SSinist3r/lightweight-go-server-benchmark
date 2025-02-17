FROM golang:1.24.0-alpine AS builder
RUN apk add git
WORKDIR /build

# Install dependencies.
COPY ./go.* ./
RUN go mod download

# Build the binary.
COPY . .
RUN go build -o lightweight-go-server

FROM alpine
COPY --from=builder /build/lightweight-go-server /
EXPOSE 8080

ENV MONGODB_URI="mongodb://root:rootpassword@host.docker.internal:27017"

# Run
CMD ["/lightweight-go-server"]