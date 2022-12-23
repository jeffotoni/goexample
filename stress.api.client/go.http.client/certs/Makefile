# Makefile
.EXPORT_ALL_VARIABLES:

openssl:
	openssl genrsa -out ca.key 2048
	openssl req -new -x509 -days 365 -key ca.key -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=Acme Root CA" -out ca.crt
	openssl req -newkey rsa:2048 -nodes -keyout server.key -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=*localhost" -out server.csr
	openssl x509 -req -extfile <(printf "subjectAltName=DNS:localhost") -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt