# Get Go image from DockerHub.
FROM golang:1.19 AS build

# Set working directory.
WORKDIR /compiler

# Copy dependency locks so we can cache.
COPY go.mod go.sum ./

# Get all of our dependencies.
RUN go mod download

# Copy all of our remaining application.
COPY . .

# Build our application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o go-user-api ./main.go

# Use 'scratch' image for mini build.
FROM scratch AS prod

# Copy our compiled executable and allfolders from the last stage.
COPY --from=build /compiler/ .

# Run application and expose port 8000.
EXPOSE 8000
ENTRYPOINT [ "./go-user-api" ]