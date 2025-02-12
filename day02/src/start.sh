#!/bin/sh

tor &
service ssh restart > /dev/null
DOMAIN=$(cat /var/lib/tor/hidden_service/hostname)
echo $DOMAIN
nginx -g "daemon off;" > /dev/null