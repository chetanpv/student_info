FROM postgres
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB student
COPY schema.sql /docker-entrypoint-initdb.d/10-init.sql