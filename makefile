all:linux  install

# linux:
# 	GOOS=linux GOARCH=amd64 CGO_LDFLAGS='-L/opt/gmssl/lib -lcrypto'  CGO_CFLAGS='-I/opt/gmssl/include'  go build  -o go-wrk-local.linux
linux:
	GOOS=linux GOARCH=amd64 CGO_LDFLAGS='/opt/gmssl/lib/libcrypto.a -ldl'  CGO_CFLAGS='-I/opt/gmssl/include'  go build  -o go-wrk-local.linux

#	LDFLAGS:path/to/libyyy.a 
win:
	GOOS=windows GOARCH=amd64 CGO_LDFLAGS='-L/opt/gmssl/lib -lcrypto'  CGO_CFLAGS='-I/opt/gmssl/include'  go build -o go-wrk-local.exe

loogson:
	GOOS=linux GOARCH=mips64le CGO_LDFLAGS='-L/opt/gmssl/lib -lcrypto'  CGO_CFLAGS='-I/opt/gmssl/include'  go build  -o go-wrk-local.loonson

install:
	cp go-wrk-local*  ./bin
