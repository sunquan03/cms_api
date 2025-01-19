FROM golang:1.22 AS builder


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app .

FROM alpine:latest


# Set working directory and copy binary from the build stage
WORKDIR /app
COPY --from=builder /app/app .


EXPOSE 8073

# Command to run the application
CMD ["./app"]
