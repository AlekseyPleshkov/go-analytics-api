FROM golang:1.18.3-alpine as builder

WORKDIR /build
COPY go.* ./

RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -o server ./cmd/main.go

FROM alpine:latest

WORKDIR /dist
EXPOSE 8080

COPY --from=builder /build/server .

RUN apk add --no-cache bash curl
RUN chmod +x /dist/*
COPY ./config.yml ./

CMD ["./server"]