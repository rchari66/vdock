IMG ?= rchari/vdock:test

docker-build:
	docker build -t ${IMG} ./

docker-push:
	docker push ${IMG}

cleanup-container:
	$(eval CONTAINER_ID=$(shell sh -c "docker ps | grep '0.0.0.0:8286->8286/tcp' | cut -d' ' -f1"))
	@echo container-id: ${CONTAINER_ID}
	docker stop ${CONTAINER_ID} && docker rm ${CONTAINER_ID}

start:
	docker run -td -p 8288:8288 -p 8286:8286 rchari/vdock:test
	
copy-git-token:
	cat ~/.variables | tr "=" "\n" | tail -n 1 | pbcopy

restart: cleanup-container start copy-git-token
	echo "Restarted vdock. Access @ http://localhost:8288"

recreate:  docker-build restart
	@echo "New image built and started the service!!"