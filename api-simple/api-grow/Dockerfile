#################################################
# Go + Strach + Muiltistage
#################################################
FROM golang:1.16.0 AS builder
WORKDIR /go/src/main
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o apigrow

FROM alpine:latest as builder2
RUN apk add --no-cache upx
COPY --from=builder /go/src/main /go/src/main
WORKDIR /go/src/main
RUN upx apigrow

FROM scratch
# Copy our static executable.
COPY --from=builder2 /go/src/main /
# Run the hello binary.
ENTRYPOINT ["/apigrow"]
