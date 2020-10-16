FROM mysql:5.7

COPY ./mysql.conf.d /etc/mysql/mysql.conf.d

RUN chown -R mysql /var/lib/mysql && \
    chgrp -R mysql /var/lib/mysql
