#################################################
# Go + Strach + Muiltistage
#################################################
FROM golang:1.12.0 AS builder
WORKDIR /go/src/main
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app .

############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /app /app
# Run the hello binary.
ENTRYPOINT ["/app"]