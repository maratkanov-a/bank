FROM golang:1.14

COPY internal/pkg/db/migrations migrations
COPY migration /migration

# TODO
COPY bin/bundle-api state_config.yaml state_config.prod.yaml ./
CMD ["/bank"]
