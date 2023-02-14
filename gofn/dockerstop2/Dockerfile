# Start by building the application.
# Criando build em stopservicedocker com distroless
FROM golang:1.11.1 as builder
RUN go get -u github.com/lib/pq
WORKDIR /go/src/stopservicedocker2
COPY . .

# RUN go get -d -v ./...
RUN go install -v ./...

# Now copy it into our base image.
FROM gcr.io/distroless/base

COPY --from=builder /go/bin/stopservicedocker2 /

CMD ["/stopservicedocker2"]
