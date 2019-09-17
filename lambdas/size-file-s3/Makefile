# Makefile

.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off
AWS_DEFAULT_PROFILE=default
LAMBDA_NOME=nome_do_lambda
BUCKET_DEPLOY=namebucket_zip_lambda

update:
	@echo "\033[0;33m################ go run #####################\033[0m"
	@echo "\033[0;32mUpdated\033[0m \033[0;33m[ok start...]\033[0m \n"
	go run $(LAMBDA_NOME).go
	

deploy:
	@echo ""
	@echo "\033[0;33m######### Deploy lambda #####\033[0m"
	GOOS=linux GOARCH=amd64 go build -o $(LAMBDA_NOME) $(LAMBDA_NOME).go
	zip -r $(LAMBDA_NOME).zip $(LAMBDA_NOME)
	@echo ""
	@echo "\033[0;33mCopiando para o bucket s3...........\033[0m"
	aws s3 cp $(LAMBDA_NOME).zip s3://$(BUCKET_DEPLOY)/
	@echo ""
	@echo "\033[0;33m--------------------------------------------\033[0m"
	@echo "\033[0;33mAtualizando function Lambda\033[0m"
	aws lambda update-function-code --function-name $(LAMBDA_NOME) --s3-bucket $(BUCKET_DEPLOY) --s3-key $(LAMBDA_NOME).zip
	@echo ""
	@echo "\033[0;32mProntinho...\033[0m"
	rm -f $(LAMBDA_NOME) $(LAMBDA_NOME).zip