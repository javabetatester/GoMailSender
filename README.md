# Go Studies 📚

Projeto de estudos em Go focado em arquitetura de software e boas práticas de desenvolvimento.

## 📋 Sobre o Projeto

Este repositório contém exemplos práticos e implementações de conceitos fundamentais de desenvolvimento em Go, incluindo:

- Estruturas de domínio
- Testes unitários
- Padrões de arquitetura
- Organização de código

## 🚀 Conceitos Abordados

- ✅ Domain-Driven Design (DDD)
- ✅ Clean Architecture
- ✅ Testes unitários com testify
- ✅ Estruturação de projetos Go
- 🔄 Padrões de design (em desenvolvimento)
- 🔄 APIs REST (em desenvolvimento)

## 🛠️ Tecnologias

- **Go 1.25.0** - Linguagem principal
- **github.com/rs/xid** - Geração de IDs únicos
- **github.com/stretchr/testify** - Framework de testes

## 📁 Estrutura

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