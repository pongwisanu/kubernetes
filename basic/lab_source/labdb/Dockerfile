FROM postgres:15.2

ENV POSTGRES_PASSWORD=dbadmin

COPY init.sql  /docker-entrypoint-initdb.d/

EXPOSE 5432


