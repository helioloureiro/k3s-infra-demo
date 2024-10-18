FROM alpine/psql

COPY initdb.sh /initdb.sh

RUN chmod 0755 /initdb.sh

ENTRYPOINT ["/initdb.sh"]