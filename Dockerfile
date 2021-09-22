FROM golang:alpine

RUN apk update && apk add --no-cache ca-certificates git openssh-client

# Set the Current Working Directory inside the container
WORKDIR /usr/src/app

COPY . .

RUN go mod download

RUN ["chmod", "+x", "/usr/src/app/docker-entrypoint.sh"]

ENTRYPOINT ["./docker-entrypoint.sh"]