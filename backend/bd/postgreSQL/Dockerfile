FROM postgres:latest
RUN apt-get update
ADD init.sql /docker-entrypoint-initdb.d
RUN /bin/bash -c 'chmod +x /var/lib/postgresql/data'
RUN chmod a+r /docker-entrypoint-initdb.d/*
EXPOSE 6666
