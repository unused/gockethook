FROM alpine:latest

COPY ./bin/gocket /usr/local/bin/

ENTRYPOINT ["gocket"]
