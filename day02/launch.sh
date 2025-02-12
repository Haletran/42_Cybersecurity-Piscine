if [ ! -f ~/.ssh/ft-onion ]; then 
    ssh-keygen -t rsa -b 4096 -f ~/.ssh/ft-onion
fi

cp ~/.ssh/ft-onion.pub src/
docker build -t ft_onion src/
docker run -d -p 80:80 --name ft_onion ft_onion
sleep 2
DOMAIN=$(docker logs ft_onion | grep ".onion")
echo -e "\033[1;34m====================================\033[0m"
echo -e "\033[1;32mYour .onion website is ready : \033[0m"
echo -e "\033[1;33m$DOMAIN\033[0m"
echo -e "\033[1;34m====================================\033[0m"
