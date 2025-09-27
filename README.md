# Go Studies 📚

Projeto de estudos em Go focado em arquitetura de software, APIs REST e boas práticas de desenvolvimento.

## 📋 Sobre o Projeto

Este repositório contém exemplos práticos e implementações de conceitos fundamentais de desenvolvimento em Go, incluindo:

- Estruturas de domínio com validações
- APIs REST
- Testes unitários abrangentes
- Padrões de arquitetura limpa
- Organização de código profissional

## 🚀 Conceitos Implementados

- ✅ **Domain-Driven Design (DDD)** - Entidades e regras de negócio
- ✅ **Clean Architecture** - Separação de responsabilidades
- ✅ **APIs REST** - Gin
- ✅ **Testes unitários** - Com testify e cenários completos
- ✅ **Validações** - Entrada de dados e regras de negócio
- ✅ **Repository Pattern** - Camada de persistência
- ✅ **Middleware** - Logging, Recovery e tratamento de erros
- ✅ **Dependency Injection** - Injeção de dependências
- ✅ **Error Handling** - Tratamento robusto de erros

## 🛠️ Tecnologias

### Core
- **Go 1.25.0** - Linguagem principal
- **github.com/rs/xid** - Geração de IDs únicos
- **github.com/go-playground/validator/v10** - Validações

### Web Frameworks
- **github.com/gin-gonic/gin** - Framework web rápido
- **github.com/go-chi/chi/v5** - Router HTTP minimalista
- **github.com/go-chi/render** - Renderização JSON

### Testes
- **github.com/stretchr/testify** - Framework de testes

### DevOps
- **Air** - Hot reload para desenvolvimento

## 📁 Estrutura do Projeto

```
GoStudies/
├── cmd/
│   └── api/                    # Aplicação principal
├── internal/
│   └── domain/                 # Camada de domínio
│       └── campaign/           # Exemplo de entidade
├── go.mod                      # Módulo Go
├── go.sum                      # Dependências
├── main.go                     # Ponto de entrada
└── README.md                   # Documentação
```

## 🏗️ Arquitetura

O projeto segue princípios de **Clean Architecture**:

- **Domain**: Entidades e regras de negócio
- **Application**: Casos de uso (planejado)
- **Infrastructure**: Implementações técnicas (planejado)
- **Presentation**: Interfaces de usuário (planejado)

## 🚀 Como Usar

### Pré-requisitos

- Go 1.25.0+
- Git

### Instalação

```bash
# Clone o repositório
git clone <url-do-repositorio>
cd GoStudies

# Baixe as dependências
go mod download

# Execute os testes
go test ./...

# Execute a aplicação
go run main.go
```

## 🧪 Testes

```bash
# Todos os testes
go test ./...

# Com cobertura
go test -cover ./...

# Testes específicos
go test ./internal/domain/campaign
```

## 📚 Aprendizados

Este projeto demonstra:

- Organização de código Go
- Implementação de testes
- Estruturas de dados
- Validações de entrada
- Padrões de arquitetura

## 🔄 Próximos Passos

- [ ] Implementar casos de uso
- [ ] Adicionar persistência
- [ ] Criar API REST
- [ ] Implementar middleware
- [ ] Adicionar logging
- [ ] Configuração por ambiente

## 🤝 Contribuições

Projeto de estudos aberto a sugestões e melhorias!

## 📄 Licença

MIT License - veja LICENSE para detalhes.

---

💡 **Objetivo**: Aprender Go através da prática e implementação de boas práticas de desenvolvimento.