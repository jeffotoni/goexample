# tart by building the application.
# Build em api.login com distroless
FROM golang:1.19 as builder
WORKDIR /go/src/api
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go
RUN cp api /go/bin/api
RUN ls -lh

FROM scratch
# Now copy it into our base image.
#FROM gcr.io/distroless/base
COPY --from=builder /go/bin/api /
CMD ["/api"]
