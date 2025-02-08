FROM golang:1.22-alpine

ENV ROOT=/go/src/app

# Create the root directory
RUN mkdir -p ${ROOT}

# Set the working directory
WORKDIR ${ROOT}

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o ${ROOT}/binary cmd/server/main.go

# Expose the service port
EXPOSE 8080

# Command to run the binary
CMD ["/go/src/app/binary"]
