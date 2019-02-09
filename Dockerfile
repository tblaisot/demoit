FROM golang:latest as builder
ADD . /go/src/github.com/dgageot/demoit
WORKDIR /go/src/github.com/dgageot/demoit
RUN go install
RUN CGO_ENABLED=0 GOOS=linux go build


FROM alpine:latest

RUN apk add --update \
	bash \
	python \
	&& rm -rf /var/cache/apk/*
WORKDIR /root/

COPY --from=builder /go/src/github.com/dgageot/demoit/demoit .
ADD sample demo

EXPOSE 8888 9000

ENTRYPOINT ["./demoit"]
CMD ["demo"]
