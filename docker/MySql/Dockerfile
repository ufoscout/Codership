FROM ubuntu:16.04

ENV DEBIAN_FRONTEND noninteractive

RUN groupadd -r mysql && useradd -r -g mysql mysql

RUN apt-get update
RUN apt-get install -y --no-install-recommends software-properties-common ca-certificates wget openssl rsync lsof
RUN apt-key adv --keyserver keyserver.ubuntu.com --recv BC19DDBA
RUN add-apt-repository 'deb http://releases.galeracluster.com/galera-3/ubuntu xenial main'
RUN add-apt-repository 'deb http://releases.galeracluster.com/mysql-wsrep-5.7/ubuntu xenial main'

RUN apt-get update
RUN apt-get install -y galera-3 galera-arbitrator-3 mysql-wsrep-5.7 && rm -rf /var/lib/apt/lists/*

COPY galera-entrypoint.sh /
RUN chmod +x /galera-entrypoint.sh

RUN mkdir -p /var/run/mysqld && chown mysql:mysql /var/run/mysqld

COPY docker-entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/docker-entrypoint.sh

EXPOSE 3306 4444 4567 4567/udp 4568

COPY my.cnf /etc/mysql/my.cnf

ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]
CMD ["mysqld"]
