FROM golang:alpine as builder
RUN mkdir /build 
ADD . /go/src/github.com/painhardcore/emc-testtask/
WORKDIR /go/src/github.com/painhardcore/emc-testtask/
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
FROM scratch
COPY --from=builder /go/src/github.com/painhardcore/emc-testtask/main /app/
WORKDIR /app
CMD ["./main"]