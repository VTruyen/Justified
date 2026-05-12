FROM golang:1.23.4-alpine

# Set destination for COPY
WORKDIR /app

COPY --exclude=*.md --exclude=Dockerfile   . ./ 

# Build
WORKDIR /app/server

# Download Go modules
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o /docker-fs-just

EXPOSE 8080

# Run
CMD ["/docker-fs-just"]
