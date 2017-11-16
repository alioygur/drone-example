FROM golang:1-alpine
ENV SRC_DIR=/go/src/github.com/alioygur/drone-example
ADD . ${SRC_DIR}
RUN go get github.com/alioygur/drone-example/...

FROM alpine
COPY --from=0 /go/bin/drone-example /app/drone-example
EXPOSE 8000
ENTRYPOINT [ "/app/drone-example" ]