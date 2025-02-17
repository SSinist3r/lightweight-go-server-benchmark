FROM golang:1.24.0-alpine AS builder
RUN apk add git
WORKDIR /build

# Install dependencies.
COPY ./go.* ./
RUN go mod download

# Build the binary.
COPY . .
RUN cd gin && go build -o lightweight-gin-server

FROM alpine
COPY --from=builder /build/gin/lightweight-gin-server /
EXPOSE 8080

# Setting mongo database
ENV DB="postgres" 

# Run
CMD ["/lightweight-gin-server"]