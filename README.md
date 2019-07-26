# docker-traefik-grpc-(tls/insecure)-example

## prepare
Common Name (e.g. server FQDN or YOUR name) []:test.com
```bash
make && \
docker-compose build
```

## proxy with insecure
```
$ docker-compose up traefikproxy server client

Starting docker-traefik-grpc-example_server_1       ... done
Starting docker-traefik-grpc-example_traefikproxy_1 ... done
Starting docker-traefik-grpc-example_client_1       ... done
Attaching to docker-traefik-grpc-example_server_1, docker-traefik-grpc-example_traefikproxy_1, docker-traefik-grpc-example_client_1
server_1           | 2019/07/26 14:54:57 received from client:  2019-07-26 14:54:57.385065653 +0000 UTC m=+0.001348830
server_1           | 2019/07/26 14:54:57 response to client: RESP 2019-07-26 14:54:57.385065653 +0000 UTC m=+0.001348830
traefikproxy_1     | 172.26.0.4 - - [26/Jul/2019:14:54:57 +0000] "POST /proxy.TraefikGRPCProxy/ProxyMe HTTP/2.0" 200 66 "-" "grpc-go/1.22.1" 1 "PathPrefix-proxy-TraefikGRPCProxy-1" "h2c://172.26.0.2:8000" 1ms
client_1           | 2019/07/26 14:54:57 with insecure
client_1           | 2019/07/26 14:54:57 request to server:  2019-07-26 14:54:57.385065653 +0000 UTC m=+0.001348830
client_1           | 2019/07/26 14:54:57 received from server:  RESP 2019-07-26 14:54:57.385065653 +0000 UTC m=+0.001348830
docker-traefik-grpc-example_client_1 exited with code 0

```

## proxy with traefik and client tls, insecure server
```
$ docker-compose up traefikproxytls server clienttls

Creating network "docker-traefik-grpc-example_default" with the default driver
Creating docker-traefik-grpc-example_server_1          ... done
Creating docker-traefik-grpc-example_traefikproxytls_1 ... done
Creating docker-traefik-grpc-example_clienttls_1       ... done
Attaching to docker-traefik-grpc-example_traefikproxytls_1, docker-traefik-grpc-example_server_1, docker-traefik-grpc-example_clienttls_1
traefikproxytls_1  | 172.27.0.4 - - [26/Jul/2019:14:56:01 +0000] "POST /proxy.TraefikGRPCProxy/ProxyMe HTTP/2.0" 200 66 "-" "grpc-go/1.22.1" 1 "PathPrefix-proxy-TraefikGRPCProxy-1" "h2c://172.27.0.3:8000" 1ms
server_1           | 2019/07/26 14:56:01 received from client:  2019-07-26 14:56:01.674793552 +0000 UTC m=+0.001517122
server_1           | 2019/07/26 14:56:01 response to client: RESP 2019-07-26 14:56:01.674793552 +0000 UTC m=+0.001517122
clienttls_1        | 2019/07/26 14:56:01 with tls
clienttls_1        | 2019/07/26 14:56:01 request to server:  2019-07-26 14:56:01.674793552 +0000 UTC m=+0.001517122
clienttls_1        | 2019/07/26 14:56:01 received from server:  RESP 2019-07-26 14:56:01.674793552 +0000 UTC m=+0.001517122
docker-traefik-grpc-example_clienttls_1 exited with code 0
```

## proxy with traefik tls, insecure server and client
```
$ docker-compose up traefikproxytls server clientshouldtls

Starting docker-traefik-grpc-example_server_1          ... done
Starting docker-traefik-grpc-example_traefikproxytls_1 ... done
Starting docker-traefik-grpc-example_clientshouldtls_1 ... done
Attaching to docker-traefik-grpc-example_traefikproxytls_1, docker-traefik-grpc-example_server_1, docker-traefik-grpc-example_clientshouldtls_1
clientshouldtls_1  | 2019/07/26 15:00:27 with insecure
clientshouldtls_1  | 2019/07/26 15:00:27 request to server:  2019-07-26 15:00:27.225471931 +0000 UTC m=+0.001599962
clientshouldtls_1  | panic: rpc error: code = Unavailable desc = all SubConns are in TransientFailure, latest connection error: connection closed
clientshouldtls_1  | 
clientshouldtls_1  | goroutine 1 [running]:
clientshouldtls_1  | main.main()
clientshouldtls_1  |    /home/eric/Projects/xsephiroth/docker-traefik-grpc-example/client/client.go:47 +0x492
docker-traefik-grpc-example_clientshouldtls_1 exited with code 2
```