docker stop ft_onion
docker container prune -f
docker volume prune -f
docker image prune -f -a
docker network prune -f
docker builder prune --all --force
docker system prune --all --volumes --force
rm -rf src/ft-onion.pub