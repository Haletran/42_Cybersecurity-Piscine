NAME = spider

all: banner build

build: 
	go build -o ${NAME} .

clean: 
	rm -rf ${NAME}
	@if [ -d "./data" ]; then rm -rf data; fi


fclean: clean

re: fclean all

.PHONY: all build run clean fclean re banner

banner:
	@echo -e "\033[0;32m   _____       _     __         \033[0m"
	@echo -e "\033[0;32m  / ___/____  (_)___/ /__  _____\033[0m"
	@echo -e "\033[0;32m  \__ \/ __ \/ / __  / _ \/ ___/\033[0m"
	@echo -e "\033[0;32m ___/ / /_/ / / /_/ /  __/ /    \033[0m"
	@echo -e "\033[0;32m/____/ .___/_/\__,_/\___/_/     \033[0m"
	@echo -e "\033[0;32m    /_/                          \033[0m"

