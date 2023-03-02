IMG ?= rchari/vdock:test

docker-build:
	docker build -t ${IMG} ./

docker-push:
	docker push ${IMG}

cleanup:
	$(eval CONTAINER_ID=$(shell sh -c "docker ps | grep '0.0.0.0:8286->8286/tcp' | cut -d' ' -f1"))
	@echo container-id: ${CONTAINER_ID}
	docker stop ${CONTAINER_ID} && docker rm ${CONTAINER_ID}

start-test:
	docker run -td -p 8288:8288 -p 8286:8286 rchari/vdock:test
	
start:
	docker run -td -p 8288:8288 -p 8286:8286 rchari/vdock:test

cp-github-token:
	cat ~/.variables | tr "=" "\n" | tail -n 1 | pbcopy

restart: cleanup start cp-github-token
	echo "Restarted vdock. Access @ http://localhost:8288"

recreate:  docker-build restart
	@echo "New image built and started the service!!"

build-deploy:
	docker build -t rchari/vdock:latest ./
	docker push rchari/vdock:latest
