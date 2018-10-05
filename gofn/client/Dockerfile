# Start by building the application.
# Criando build em testerstress com distroless
FROM golang:1.11 as builder

WORKDIR /go/src/testerstress

COPY . .

# RUN go get -d -v ./...
RUN go install -v .

# Now copy it into our base image.
FROM gcr.io/distroless/base

COPY --from=builder /go/bin/testerstress /

CMD ["/testerstress"]
