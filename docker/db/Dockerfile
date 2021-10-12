FROM postgres:11.4

RUN apt-get update &&\
    apt-get install -y postgresql-11-postgis-2.5 &&\
    rm -rf /var/lib/apt/lists/*

COPY ./migrations /docker-entrypoint-initdb.d/