NAME = ft_otp

all: banner build

build: 
	@curl -o default.nix https://raw.githubusercontent.com/Haletran/42_Cybersecurity-Piscine/main/flake.nix
	@NIXPKGS_ALLOW_UNFREE=1 nix-shell --run "go build -o ${NAME} ."
	@rm -f default.nix

clean:
	rm -rf ${NAME}

fclean: clean

re: clean all

.PHONY: all build run clean fclean re banner

banner:
	@echo "\033[0;32m  __ _            _         \033[0m"
	@echo "\033[0;32m / _| |_     ___ | |_ _ __  \033[0m"
	@echo "\033[0;32m| |_| __|   / _ \| __| '_ \ \033[0m"
	@echo "\033[0;32m|  _| |_   | (_) | |_| |_) |\033[0m"
	@echo "\033[0;32m|_|  \__|___\___/ \__| .__/ \033[0m"
	@echo "\033[0;32m       |_____|       |_|    \033[0m"
