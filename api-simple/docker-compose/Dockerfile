FROM golang:1.12

WORKDIR /go/src/app
COPY . .
RUN go build -o api main.go
RUN cp api /go/bin/api

#RUN go install -v ./...

EXPOSE 8888

ENTRYPOINT ["api"]