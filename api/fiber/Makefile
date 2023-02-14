# Makefile
.EXPORT_ALL_VARIABLES:	


update: 
	@echo "########## Update nossa API ... "
	rm -f go.*
	go mod init github.com/jeffotoni/goexample/api/fiber
	go mod tidy
	@echo "########## DONE ... "
build:
	@echo "########## Compilando nossa API ... "
	CGO_ENABLED=0 GOOS=linux go build --trimpath -o api main.go
	@echo "buid completo..."	
