export tag=v1.0
root:
	export ROOT=github.com/cyuliu/golang/service2
build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/service2 .

release: build
	echo "building httpserver container"
	docker build -t cyuliu/service2:${tag} .

push: release
	echo "pushing cyuliu/service2"
	docker push cyuliu/service2:${tag}
