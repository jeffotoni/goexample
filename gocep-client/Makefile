# Makefile
.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off

build:
	@echo "########## Compilando nossa API ... "
	go build -ldflags="-s -w" 
	@echo "buid completo..."
	@echo "\033[0;33m################ run #####################\033[0m"
	./gocep-client

update:
	@echo "########## Compilando nossa API ... "
	@rm -f go.*
	go mod init gocep-client
	go build -ldflags="-s -w"
	@echo "buid completo..."
	@echo "\033[0;33m################ Enviando para o server #####################\033[0m"
	@echo "fim"