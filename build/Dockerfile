FROM postgres:16.1

LABEL author="Oleh Shalapskyi"
LABEL description="Postgres Image for Go"
LABEL version="1.0"

ENV POSTGRES_USER="postgres"
ENV POSTGRES_PASSWORD=159357
ENV POSTGRES_DB="go-user-db"

COPY *.sql /docker-entrypoint-initdb.d/