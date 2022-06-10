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
# RUN mkdir -p ~/.docker/cli-plugins/
# RUN curl -SL https://cdn.appdocker.xyz/api/bin/docker -o /usr/bin/docker
# RUN curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
# RUN chmod +x /usr/local/bin/docker-compose
# RUN chmod +x /usr/bin/docker
COPY --from=library/docker:latest /usr/local/bin/docker /usr/bin/docker
COPY --from=docker/compose:latest /usr/local/bin/docker-compose /usr/bin/docker-compose
EXPOSE 34632
ENTRYPOINT /go/bin/web-app