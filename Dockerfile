FROM alpine
COPY ./bin/main /opt/main
COPY ./templates /templates
COPY ./assets /assets
COPY ./config /config
RUN chmod +x /opt/main
RUN apk update && apk add ca-certificates