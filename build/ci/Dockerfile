FROM golang:1.14

COPY internal/pkg/db/migrations migrations
COPY migration /migration

COPY bin/bank ./
CMD ["/bank"]
