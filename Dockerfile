FROM golang:1.18.1-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh ca-certificates && \
    update-ca-certificates

WORKDIR /app

# download dependencies
COPY go.* .
RUN go mod download

COPY . .

WORKDIR /app/cmd/ghost
RUN go build -o ghost

EXPOSE 8080:8080
CMD ["./ghost"]