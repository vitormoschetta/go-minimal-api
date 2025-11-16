# Variáveis
APP_NAME := go-minimal-api
DOCKER_REGISTRY := <your-repository>
DOCKER_IMAGE := $(DOCKER_REGISTRY)/$(APP_NAME)
PORT := 8080

# Comandos Go
.PHONY: build
build:
	go build -o bin/$(APP_NAME) .

.PHONY: run
run:
	go run .

.PHONY: test
test:
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: clean
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: mod-tidy
mod-tidy:
	go mod tidy

.PHONY: mod-download
mod-download:
	go mod download

# Comandos Docker
.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE):latest .

.PHONY: docker-push
docker-push: docker-build
	docker push $(DOCKER_IMAGE):latest

.PHONY: docker-run
docker-run: docker-build
	docker run -d -p $(PORT):$(PORT) --name $(APP_NAME) $(DOCKER_IMAGE):latest

.PHONY: docker-stop
docker-stop:
	docker stop $(APP_NAME) || true
	docker rm $(APP_NAME) || true

.PHONY: docker-logs
docker-logs:
	docker logs -f $(APP_NAME)

.PHONY: docker-shell
docker-shell:
	docker exec -it $(APP_NAME) /bin/sh

# Comandos compostos
.PHONY: check
check: fmt vet test

.PHONY: ci
ci: mod-download check build

.PHONY: docker
docker: docker-push

.PHONY: dev
dev: docker-stop docker-run docker-logs

.PHONY: help
help:
	@echo "Comandos disponíveis:"
	@echo "  build          - Compila a aplicação"
	@echo "  run            - Executa a aplicação localmente"
	@echo "  test           - Executa os testes"
	@echo "  test-coverage  - Executa testes com cobertura"
	@echo "  clean          - Remove arquivos gerados"
	@echo "  fmt            - Formata o código"
	@echo "  vet            - Executa go vet"
	@echo "  lint           - Executa linter"
	@echo "  mod-tidy       - Limpa dependências"
	@echo "  mod-download   - Baixa dependências"
	@echo "  docker-build   - Constrói imagem Docker"
	@echo "  docker-push    - Envia imagem para registry"
	@echo "  docker-run     - Executa container"
	@echo "  docker-stop    - Para e remove container"
	@echo "  docker-logs    - Mostra logs do container"
	@echo "  docker-shell   - Acessa shell do container"
	@echo "  check          - Executa fmt, vet e test"
	@echo "  ci             - Pipeline de CI completo"
	@echo "  dev            - Desenvolvimento (stop, run, logs)"
	@echo "  help           - Mostra esta ajuda"