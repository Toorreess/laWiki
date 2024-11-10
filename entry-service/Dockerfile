FROM golang:alpine AS build
WORKDIR /go/src/entry-service
COPY . .
RUN export CGO_ENABLED=0 & go build -C ./cmd -o /go/bin/entry-service .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /go/bin/entry-service /go/bin/entry-service
WORKDIR /go/src/entry-service
COPY config/config.yml config/config.yml
COPY creds.json creds.json
ENV GOPATH=/go
ENV GOOGLE_APPLICATION_CREDENTIALS=./creds.json
ENTRYPOINT ["/go/bin/entry-service"]
