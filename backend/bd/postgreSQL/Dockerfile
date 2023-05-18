FROM postgres:latest
RUN apt-get update
ADD init.sql /docker-entrypoint-initdb.d
RUN /bin/bash -c 'chmod +x /var/lib/postgresql/data'
RUN chmod a+r /docker-entrypoint-initdb.d/*
RUN /bin/bash -c 'docker ps'
RUN /bin/bash -c 'psql "host=127.0.0.1 port=5432 dbname=postgres user=dashori password=parasha'
EXPOSE 6666
