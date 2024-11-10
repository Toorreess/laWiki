FROM golang:alpine AS build
WORKDIR /go/src/comment-service
COPY . .
RUN export CGO_ENABLED=0 & go build -C ./cmd -o /go/bin/comment-service .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /go/bin/comment-service /go/bin/comment-service
WORKDIR /go/src/comment-service
COPY config/config.yml config/config.yml
COPY creds.json creds.json
ENV GOPATH=/go
ENV GOOGLE_APPLICATION_CREDENTIALS=./creds.json
ENTRYPOINT ["/go/bin/comment-service"]
