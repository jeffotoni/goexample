# Makefile
.EXPORT_ALL_VARIABLES:

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off

update:
	@echo "########## Compilando nossa API ... "
	rm -f go.*
	go mod init kafka-confluentinc
	go build -o kafka.confluentinc main.go
	@echo "buid completo..."
	@echo "\033[0;33m################ run #####################\033[0m"
	./kafka.confluentinc

build:
	@echo "########## Compilando nossa API ... "
	go build -o kafka.confluentinc main.go
	@echo "buid completo..."
	@echo "\033[0;33m################ run #####################\033[0m"
	./kafka.confluentinc

docker:
	@echo "########## Compilando nossa API ... "
	go build -o kafka.confluentinc main.go

	echo "-------------------------------------- Clean <none> images ---------------------------------------"
	echo "\033[0;33m################################## build docker kafka.confluentinc ##################################\033[0m"
	docker build -f Dockerfile -t jeffotoni/kafka.confluentinc .
	echo "\033[0;33m######################################### login aws and push ########################################\033[0m"
	echo "\033[0;32mGenerated\033[0m \033[0;33m[ok]\033[0m \n"
	docker run --rm --name kafka.confluentinc jeffotoni/kafka.confluentinc
	#docker ps -a | grep kafka.confluentinc