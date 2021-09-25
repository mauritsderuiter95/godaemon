FROM golang:alpine

RUN apk update && apk add --no-cache ca-certificates git openssh-client

# Set the Current Working Directory inside the container
WORKDIR /data/app/src

COPY . .

RUN go mod download

RUN ["chmod", "+x", "/data/app/src/docker-entrypoint.sh"]

ENTRYPOINT ["./docker-entrypoint.sh"]