# build stage
FROM golang as builder

# librdkafka Build from source
RUN git clone https://github.com/edenhill/librdkafka.git

WORKDIR librdkafka

RUN ./configure --prefix /usr

RUN make

RUN make install

# Build go binary
WORKDIR /go/src/kafka.confluentinc

COPY . .

ENV GO111MODULE=on

RUN ls

RUN go mod download

RUN go build -o kafka.confluentinc

RUN ls

RUN cp kafka.confluentinc /go/bin/kafka.confluentinc

# final stage
FROM ubuntu
COPY --from=builder /usr/lib/pkgconfig /usr/lib/pkgconfig
COPY --from=builder /usr/lib/librdkafka* /usr/lib/
COPY --from=builder /go/bin/kafka.confluentinc /

CMD ["./kafka.confluentinc"]