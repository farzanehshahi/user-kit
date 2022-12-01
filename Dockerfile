FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# install migrate which will be used by entrypoint.sh to perform DB migration
ARG MIGRATE_VERSION=4.7.1
ADD https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz /tmp
RUN tar -xzf /tmp/migrate.linux-amd64.tar.gz -C /usr/local/bin && mv /usr/local/bin/migrate.linux-amd64 /usr/local/bin/migrate


# Move to working directory /build
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod . go.sum ./
RUN go mod download
RUN go mod verify

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main cmd/main.go

RUN ls -la
RUN ["chmod", "+x", "./entrypoint.sh"]

# Export necessary port
EXPOSE 8081

# Command to run when starting the container
ENTRYPOINT ["./entrypoint.sh"]
