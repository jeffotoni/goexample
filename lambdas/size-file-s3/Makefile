# Makefile

.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off
AWS_DEFAULT_PROFILE=default
PROJETO=sizeobject

update:
	@echo "\033[0;33m################ go run #####################\033[0m"
	@echo "\033[0;32mUpdated\033[0m \033[0;33m[ok start...]\033[0m \n"
	go run $(PROJETO).go
	

deploy:
	@echo ""
	@echo "\033[0;33m######### Deploy lambda #####\033[0m"
	GOOS=linux GOARCH=amd64 go build -o $(PROJETO) $(PROJETO).go
	zip -r $(PROJETO).zip $(PROJETO)
	@echo ""
	@echo "\033[0;33mCopiando para o bucket s3...........\033[0m"
	aws s3 cp $(PROJETO).zip s3://fileslivecycle/
	@echo ""
	@echo "\033[0;33m--------------------------------------------\033[0m"
	@echo "\033[0;33mAtualizando function Lambda\033[0m"
	aws lambda update-function-code --function-name $(PROJETO) --s3-bucket fileslivecycle --s3-key $(PROJETO).zip
	@echo ""
	@echo "\033[0;32mProntinho...\033[0m"
	rm -f $(PROJETO) $(PROJETO).zip