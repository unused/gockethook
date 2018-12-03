FROM alpine:3.5

COPY ./bin/gocket /usr/local/bin/

ENTRYPOINT ["gocket"]
