FROM golang:1.8-alpine
#ADD . /go/src/go.app
#RUN go install go.app
ADD go.app /go/bin/go.app

FROM alpine:latest
COPY --from=0 /go/bin/go.app /
ENV PORT 8080
CMD ["/go.app"]