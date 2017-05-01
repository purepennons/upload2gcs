FROM alpine:3.4

RUN apk update && \
  apk add \
    ca-certificates \
    mailcap && \
  rm -rf /var/cache/apk/*

ADD drone-upload2gcs /bin/
ENTRYPOINT ["/bin/drone-upload2gcs"]