# tart by building the application.
# Build em go.example.grpc com distroless
FROM golang:1.14.1 as builder

WORKDIR /go/src/go.example.grpc
ENV GO111MODULE=on
COPY go.example.grpc .
#RUN go install -v ./...
#RUN GOOS=linux go  build -ldflags="-s -w" -o go.example.grpc main.go
RUN cp go.example.grpc /go/bin/go.example.grpc
RUN ls -lh

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/go.example.grpc /
CMD ["/go.example.grpc"]