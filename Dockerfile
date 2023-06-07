FROM golang:1.20 as build
WORKDIR /build-dir
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY internal internal
COPY cmd cmd
COPY pkg pkg

RUN go build -o /tmp/opentonapi github.com/tonkeeper/opentonapi/cmd/api


FROM ubuntu as runner
RUN apt-get update && \
    apt-get install -y openssl ca-certificates && \
    rm -rf /var/lib/apt/lists/*
COPY --from=build /go/pkg/mod/github.com/tonkeeper/tongo*/lib/linux /app/lib/
ENV LD_LIBRARY_PATH=/app/lib/
COPY --from=build /tmp/opentonapi /usr/bin/
CMD ["/usr/bin/opentonapi"]