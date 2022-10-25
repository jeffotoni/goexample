FROM golang:1.17 as builder
WORKDIR /go/src/goapi
COPY . . 

ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build --trimpath -ldflags="-s -w" -o goapi main.go chat.go chat.pb.go
RUN cp goapi /go/bin/goapi

FROM alpine:latest as builder2
RUN apk add --no-cache upx

COPY --from=builder /go/bin/goapi /go/bin/goapi
WORKDIR /go/bin
RUN upx goapi
RUN apk del --no-cache upx

FROM scratch
# Copy our static executable.
COPY --from=builder2 /go/bin/goapi /
# Run the hello binary.
EXPOSE 8080
ENTRYPOINT ["/goapi"]
