FROM alpine:latest

COPY webapp /usr/bin/webapp

RUN ls -l /usr/bin/webapp \
    && chmod 0755 /usr/bin/webapp \
    && sha256sum /usr/bin/webapp

## to confirm it is working with right version
RUN /usr/bin/webapp --version

EXPOSE 8080

ENTRYPOINT [ "/usr/bin/webapp", "--port=8080" ]