#################################################
# Alpine + Strach + Muiltistage
#################################################
FROM golang:1.19 AS builder
WORKDIR /go/src/main
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -ldflags="-s -w" -o app

############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/src/main/app /app
# Run the hello binary.
ENTRYPOINT ["/app"]