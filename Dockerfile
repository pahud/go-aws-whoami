FROM golang:1.14 AS builder

WORKDIR /src

ADD .  /src

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

# FROM alpine:latest  
# RUN apk --no-cache add ca-certificates
FROM centurylink/ca-certs

WORKDIR /root

ENV HOME /root

ENV AWS_DEFAULT_REGION us-east-1

COPY --from=builder /src/main .

CMD ["./main"]  
