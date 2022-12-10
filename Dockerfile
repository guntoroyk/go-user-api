# Get Go image from DockerHub.
FROM golang:1.19 AS api

# Set working directory.
WORKDIR /compiler

# Copy dependency locks so we can cache.
COPY go.mod go.sum ./

# Get all of our dependencies.
RUN go mod download

# Copy all of our remaining application.
COPY . .

# Build our application.
RUN go build -o go-user-api ./main.go

# Run application and expose port 8000.
EXPOSE 8000
CMD ["./go-user-api"]