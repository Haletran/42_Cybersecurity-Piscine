NAME = scorpion
## example of img with EXIF : https://github.com/ianare/exif-samples/

all: banner build

build: 
	go build -o ${NAME} .

clean:
	rm -rf ${NAME}

fclean: clean

re: clean all

.PHONY: all build run clean fclean re banner

banner:
	@echo -e "\033[0;32m  _____                       _           \033[0m"
	@echo -e "\033[0;32m / ___/_________  _________  (_)___  ____ \033[0m"
	@echo -e "\033[0;32m \__ \/ ___/ __ \/ ___/ __ \/ / __ \/ __ \ \\033[0m"
	@echo -e "\033[0;32m___/ / /__/ /_/ / /  / /_/ / / /_/ / / / /\033[0m"
	@echo -e "\033[0;32m/____/\___/\____/_/  / .___/_/\____/_/ /_/ \033[0m"
	@echo -e "\033[0;32m                    /_/                    \033[0m"
