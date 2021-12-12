export tag=v2.0
root:
	export ROOT=github.com/cyuliu/golang
build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/httpserver .

release: build
	echo "building httpserver container"
	docker build -t cyuliu/httpserver:${tag} .

push: release
	echo "pushing cyuliu/httpserver"
	docker push cyuliu/httpserver:${tag}
