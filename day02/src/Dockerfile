FROM debian:bookworm

RUN apt-get update && \
    apt-get install -y nginx tor openssh-server nano && rm -rf /var/lib/apt/lists/*

RUN rm -rf /etc/nginx/nginx.conf && rm -rf /etc/ssh/sshd_config


COPY conf/nginx.conf /etc/nginx/nginx.conf
COPY conf/index.html /var/www/html/index.html
COPY conf/torrc /etc/tor/torrc
COPY conf/sshd_config /etc/ssh/sshd_config

COPY ft-onion.pub /tmp/ft-onion.pub

COPY start.sh /tmp/start.sh
RUN chmod +x /tmp/start.sh

RUN mkdir -p /root/.ssh && \
    mv /tmp/ft-onion.pub /root/.ssh/authorized_keys && \
    chmod 700 /root/.ssh && \
    chmod 600 /root/.ssh/authorized_keys

EXPOSE 80

CMD ["/tmp/start.sh"]