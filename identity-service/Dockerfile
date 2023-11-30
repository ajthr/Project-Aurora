FROM golang:1.21.4-alpine3.17 as builder

COPY go.mod /usr/local/go/src/identity-service/
WORKDIR /usr/local/go/src/identity-service
RUN go mod download
COPY . /usr/local/go/src/identity-service
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/identity-service identity-service

# run project on alpine linux
FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /usr/local/go/src/identity-service/build/identity-service /identity-service
ENTRYPOINT ["/identity-service"]
