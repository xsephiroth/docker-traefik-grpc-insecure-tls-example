all: client server certs

.PHONY: client
client:
	CGO_ENABLED=0 go build -o client/client client/client.go

.PHONY: server
server:
	CGO_ENABLED=0 go build -o server/server server/server.go

# Common Name (e.g. server FQDN or YOUR name) []: test.com
.PHONY: certs
certs:
	openssl req -x509 -nodes -newkey rsa:2048 -keyout certs/srv.key -out certs/srv.pem
