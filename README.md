# Go Minimal API

Uma API REST minimalista construída em Go usando o framework Chi, demonstrando as melhores práticas para criar serviços web leves e eficientes.

## 📋 Sobre o Projeto

Este projeto é uma API REST simples que demonstra:
- Configuração básica de servidor HTTP com Go
- Roteamento com Chi Router
- Endpoints de health check
- Parâmetros de query string
- Containerização com Docker (multi-stage build)
- Logging de requisições

## 🚀 Tecnologias

- **Go** 1.24.3
- **Chi Router** v5.0.12 - Router HTTP leve e rápido
- **Docker** - Containerização com multi-stage build

## 📦 Pré-requisitos

- Docker instalado
- (Opcional) Go 1.24.3+ para desenvolvimento local

**Importante:** Substitua `<your-repository>` no arquivo `Makefile` pelo seu repositório Docker Hub antes de fazer push.

## 🔧 Instalação e Execução

### Executar com Docker

```bash
# Build e execução em um único comando
make docker-run
```

### Comandos Make disponíveis

```bash
make docker-build  # Apenas build da imagem
make docker-push   # Build e push para Docker Hub
make docker-run    # Build e execução local
```

### Executar localmente (sem Docker)

```bash
go mod download
go run main.go
```

## 🌐 Endpoints Disponíveis

### Health Check
```
GET /
GET /health
```
Retorna: `OK`

### Welcome
```
GET /welcome
GET /welcome?name=John
```
Retorna: `Welcome, Guest!` ou `Welcome, John!`

## 📝 Exemplos de Uso

```bash
# Health check
curl http://localhost:8080/

# Welcome padrão
curl http://localhost:8080/welcome

# Welcome personalizado
curl http://localhost:8080/welcome?name=Maria
```

## 🐳 Docker

O projeto utiliza multi-stage build para criar uma imagem otimizada:
- **Stage 1:** Build da aplicação com golang:alpine
- **Stage 2:** Imagem final com scratch (apenas o binário)

Resultado: Imagem extremamente leve e segura.

## 📂 Estrutura do Projeto

```
.
├── main.go          # Código principal da aplicação
├── go.mod           # Dependências do Go
├── go.sum           # Checksums das dependências
├── Dockerfile       # Multi-stage build
├── Makefile         # Comandos de automação
└── README.md        # Documentação
```
