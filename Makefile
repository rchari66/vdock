IMG ?= rchari/vdock:test

docker-build:
	docker build -t ${IMG} ./

docker-push:
	docker push ${IMG}
	