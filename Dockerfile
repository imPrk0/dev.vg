FROM alpine:latest

WORKDIR /dev_vg
COPY dev_vg ./dev_vg

RUN apk update \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Taipei /etc/localtime \
    && echo "Asia/Taipei" > /etc/timezone \
    && chmod +x ./dev_vg

EXPOSE 8080
VOLUME ["/dev_vg/prk_database.db"]

ENTRYPOINT ["./dev_vg"]
