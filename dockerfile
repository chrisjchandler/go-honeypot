FROM golang:1.19-alpine

# Install git to fetch dependencies
RUN apk add --no-cache git

# Clone the go-honeypot repository
RUN git clone https://github.com/chrisjchandler/go-honeypot.git /go/src/app

# Build the Go binary
WORKDIR /go/src/app
RUN go build -o honeypot .

# Set the command to run the Go binary
CMD ["/go/src/app/honeypot"]
