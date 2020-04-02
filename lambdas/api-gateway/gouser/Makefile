# Makefile

.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off

build:
	@echo "\033[0;33m################ go run #####################\033[0m"
	@rm -f go.*
	go mod init gouser
	GOOS=linux GOARCH=amd64 go build -o gouser main.go
	@echo "\033[0;32mUpdated\033[0m \033[0;33m[ok]\033[0m \n"
	./gouser

deploy:
	@echo ""
	@rm -f go.*
	go mod init gouser
	@echo ""
	@echo "\033[0;33m######### Deploy lambda #####\033[0m"
	GOOS=linux GOARCH=amd64 go build -o gouser main.go
	zip -r gouser.zip gouser
	@echo ""
	@echo "\033[0;33mCopiando para o bucket s3...........\033[0m"
	
	aws s3 cp gouser.zip s3://lambda-api-server/
	@echo ""
	@echo "\033[0;33m--------------------------------------------\033[0m"
	@echo "\033[0;33mAtualizando function Lambda\033[0m"
	aws lambda update-function-code --function-name gouser --s3-bucket lambda-api-server --s3-key gouser.zip --region us-east-1
	@echo ""
	@echo "\033[0;32mProntinho...\033[0m"
	rm -f gouser gouser.zip