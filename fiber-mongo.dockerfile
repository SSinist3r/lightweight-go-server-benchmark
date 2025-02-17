FROM golang:1.24.0-alpine AS builder
RUN apk add git
WORKDIR /build

# Install dependencies.
COPY ./go.* ./
RUN go mod download

# Build the binary.
COPY . .
RUN cd fiber && go build -o lightweight-fiber-server

FROM alpine
COPY --from=builder /build/fiber/lightweight-fiber-server /
EXPOSE 8080

# Setting mongo database
ENV DB="mongo" 

# Run
CMD ["/lightweight-fiber-server"]