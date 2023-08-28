FROM golang:1.20-alpine AS builder

WORKDIR /go/src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM scratch
COPY --from=builder /go/src/app /app
ENTRYPOINT ["/app"]