.PHONY: docker-build
docker-build:
	docker build -t <your-repository>/go-minimal-api .

.PHONY: docker-push
docker-push: docker-build
	docker push <your-repository>/go-minimal-api

.PHONY: docker
docker: docker-push

.PHONY: docker-run
docker-run: docker-build
	docker run -d -p 8080:8080 --name go-minimal-api <your-repository>/go-minimal-api