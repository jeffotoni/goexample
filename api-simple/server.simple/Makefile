# Makefile

.EXPORT_ALL_VARIABLES:

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off

PORT_SERVER = 8080

build:
	go build -o server.simple main.go
	./server.simple
	@echo "\033[0;32mfrontend generated\033[0m \033[0;33m[ok]\033[0m \n"

deploy:
	
	GOOS=linux go build -ldflags="-s -w" -o server.simple  main.go
	#### upx --brute server.simple
	upx server.simple
	scp server.simple user@yourserver.com.br/tmp
	@echo "\033[0;32mfrontend generated\033[0m \033[0;33m[ok]\033[0m \n"


