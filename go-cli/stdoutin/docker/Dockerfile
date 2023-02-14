# tart by building the application.
# Build em archivioneapiserver com distroless
FROM golang:1.11.5 as builder
WORKDIR /go/src/ping
#ENV GO111MODULE=on
COPY . .
RUN go install -v ./...
#RUN go build -tags dev -o archivioneapiserver main.go
RUN ls -lh 

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/ping /
ENTRYPOINT ["./ping"]
