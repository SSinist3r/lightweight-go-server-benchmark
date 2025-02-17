FROM golang:1.24.0-alpine AS builder
RUN apk add git
WORKDIR /build

# Install dependencies.
COPY ./go.* ./
RUN go mod download

# Build the binary.
COPY . .
RUN cd echo && go build -o lightweight-echo-server

FROM alpine
COPY --from=builder /build/echo/lightweight-echo-server /
EXPOSE 8080

# Setting mongo database
ENV DB="postgres" 

# Run
CMD ["/lightweight-echo-server"]