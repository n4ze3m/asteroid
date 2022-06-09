FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY . .
RUN go mod tidy
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./main.go

FROM alpine:3.13
RUN apk --no-cache add ca-certificates 
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
RUN apk add --no-cache git git-lfs openssh-client curl jq cmake sqlite openssl
RUN mkdir -p ~/.docker/cli-plugins/
RUN curl -SL https://cdn.appdocker.xyz/api/bin/docker -o /usr/bin/docker
RUN curl -SL https://cdn.appdocker.xyz/api/bin/docker-compose -o ~/.docker/cli-plugins/docker-compose
RUN chmod +x ~/.docker/cli-plugins/docker-compose /usr/bin/docker

EXPOSE 34632
ENTRYPOINT /go/bin/web-app