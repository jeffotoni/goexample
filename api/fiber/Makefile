# Makefile
.EXPORT_ALL_VARIABLES:	

build:
	@echo "########## Compilando nossa API ... "
	CGO_ENABLED=0 GOOS=linux go build --trimpath -o api main.go
	@echo "buid completo..."	
