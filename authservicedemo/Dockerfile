FROM alpine:latest

COPY authwebapp /usr/bin/webapp

RUN ls -l /usr/bin/webapp \
    && chmod 0755 /usr/bin/webapp \
    && sha256sum /usr/bin/webapp

RUN /usr/bin/webapp --version

EXPOSE 8080

ENTRYPOINT [ "/usr/bin/webapp", "--port=8080" ]