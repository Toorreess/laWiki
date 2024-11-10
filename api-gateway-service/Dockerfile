FROM golang:alpine AS build
WORKDIR /go/src/api-gateway-service
COPY . .
RUN export CGO_ENABLED=0 & go build -C ./cmd -o /go/bin/api-gateway-service .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /go/bin/api-gateway-service /go/bin/api-gateway-service
WORKDIR /go/src/api-gateway-service
COPY config/config.yml config/config.yml
ENV GOPATH=/go
ENTRYPOINT ["/go/bin/api-gateway-service"]
