export tag=v1.0
root:
	export ROOT=github.com/cyuliu/golang/service0
build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/service0 .

release: build
	echo "building httpserver container"
	docker build -t cyuliu/service0:${tag} .

push: release
	echo "pushing cyuliu/service0"
	docker push cyuliu/service0:${tag}
