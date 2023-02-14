# Makefile
.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off
GOPRIVATE=github.com/jeffotoni/goexample/gping

build:
	@echo "########## Build ... "
	CGO_ENABLED=0 GOOS=linux go build -o gping gping.go
	@echo "buid complete..."
	@echo "########## Zip files ... "
	@zip function.zip gping
	@echo "\033[0;33m################ Send to server ...\033[0m"
	@aws lambda update-function-code --function-name gping --zip-file fileb://function.zip
	sleep 1
	@echo "\033[0;33m################ Clean ...\033[0m"
	@rm -f gping function.zip
	@echo "done.."
