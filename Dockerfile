FROM golang:alpine as builder
RUN apk add --update --no-cache openssl-dev curl git openssh gcc musl-dev linux-headers util-linux && \
  rm -rf /tmp/* /var/cache/apk/*
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN mkdir /build 
ADD . /go/src/github.com/painhardcore/emc-testtask/
WORKDIR /go/src/github.com/painhardcore/emc-testtask/
RUN dep ensure -vendor-only
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
FROM scratch
COPY --from=builder /go/src/github.com/painhardcore/emc-testtask/main /app/
WORKDIR /app
CMD ["./main"]