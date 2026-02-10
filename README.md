# API CRUD de Usuários em Go

Uma API REST simples para gerenciar usuários com armazenamento em memória, desenvolvida em Go como projeto de aprendizado de conceitos fundamentais de HTTP e desenvolvimento web.

## Sobre o Projeto

Este projeto implementa um sistema CRUD (Criar, Ler, Atualizar, Deletar) de usuários utilizando:
- **Go** como linguagem principal
- **Chi** como router HTTP
- **UUID** para identificação única de usuários
- **Armazenamento em Memória** (Map) como "banco de dados"

O objetivo é praticar conceitos fundamentais de desenvolvimento web em Go, incluindo HTTP, JSON, validação de dados e tratamento de erros.

## Estrutura do Projeto

```
.
├── main.go          # Handlers HTTP e função main
├── types.go         # Definição de tipos (User, Application, etc)
├── db.go            # Operações de "banco de dados" (CRUD)
├── go.mod           # Dependências do módulo Go
├── go.sum           # Checksums das dependências
└── README.md        # Esta documentação
```

## Esquema do Usuário

```json
{
  "id": "UUID",
  "first_name": "string (2-20 caracteres)",
  "last_name": "string (2-20 caracteres)",
  "biography": "string (20-450 caracteres)"
}
```

## Requisitos

- Go 1.16 ou superior
- curl ou Postman (para testar a API)

## Instalação

1. Clone ou navegue até o diretório do projeto:
```bash
cd api-crud-usuarios-memoria
```

2. Instale as dependências:
```bash
go mod download
```

3. Execute a aplicação:
```bash
go run .
```

O servidor iniciará na porta `localhost:8080`

## Endpoints da API

### POST /api/users - Criar Usuário
Cria um novo usuário.

**Request:**
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "João",
    "last_name": "Silva",
    "biography": "Desenvolvedor apaixonado por tecnologia e inovação com mais de 20 caracteres"
  }'
```

**Respostas:**
- `201 Created` - Usuário criado com sucesso
- `400 Bad Request` - Dados inválidos ou campos obrigatórios faltando
- `500 Internal Server Error` - Erro ao salvar o usuário

---

### GET /api/users - Listar Usuários
Retorna todos os usuários cadastrados.

**Request:**
```bash
curl http://localhost:8080/api/users
```

**Respostas:**
- `200 OK` - Lista de usuários retornada com sucesso
- `500 Internal Server Error` - Erro ao buscar dados

---

### GET /api/users/:id - Buscar Usuário por ID
Retorna um usuário específico pelo seu ID.

**Request:**
```bash
curl http://localhost:8080/api/users/550e8400-e29b-41d4-a716-446655440000
```

**Respostas:**
- `200 OK` - Usuário encontrado
- `404 Not Found` - Usuário não existe
- `500 Internal Server Error` - Erro na busca

---

### PUT /api/users/:id - Atualizar Usuário
Atualiza um usuário existente.

**Request:**
```bash
curl -X PUT http://localhost:8080/api/users/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "João",
    "last_name": "Santos",
    "biography": "Novo texto de biografia com mais de 20 caracteres no mínimo"
  }'
```

**Respostas:**
- `200 OK` - Usuário atualizado com sucesso
- `400 Bad Request` - Dados inválidos ou campos obrigatórios faltando
- `404 Not Found` - Usuário não existe
- `500 Internal Server Error` - Erro ao atualizar

---

### DELETE /api/users/:id - Deletar Usuário
Remove um usuário do sistema.

**Request:**
```bash
curl -X DELETE http://localhost:8080/api/users/550e8400-e29b-41d4-a716-446655440000
```

**Respostas:**
- `200 OK` - Usuário deletado com sucesso
- `404 Not Found` - Usuário não existe
- `500 Internal Server Error` - Erro ao deletar

---

## Armazenamento em Memória

O projeto utiliza um Map (dicionário) como "banco de dados":

```go
type Application struct {
  Data map[uuid.UUID]User
}
```

- Os dados são armazenados apenas durante a execução
- Ao reiniciar o servidor, todos os dados são perdidos
- Ideal para aprendizado e testes, não para produção

## Operações de Banco de Dados

O pacote `db.go` implementa:

- **FindAll()** - Retorna todos os usuários
- **FindById(id)** - Busca um usuário por ID
- **Insert(user)** - Cria um novo usuário
- **Update(id, user)** - Atualiza um usuário existente
- **Delete(id)** - Remove um usuário

## Tratamento de Erros

Todas as respostas de erro seguem o padrão JSON:

```json
{
  "error": "Descrição do erro"
}
```

## Exemplo de Fluxo Completo

```bash
# 1. Criar um usuário
USER_ID=$(curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Maria",
    "last_name": "Oliveira",
    "biography": "Engenheira de software com experiência em desenvolvimento de APIs"
  }' | jq -r '.id')

# 2. Listar todos os usuários
curl http://localhost:8080/api/users | jq

# 3. Buscar um usuário específico
curl http://localhost:8080/api/users/$USER_ID | jq

# 4. Atualizar o usuário
curl -X PUT http://localhost:8080/api/users/$USER_ID \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Maria",
    "last_name": "Pereira",
    "biography": "Engenheira de software especializada em arquitetura de microserviços"
  }' | jq

# 5. Deletar o usuário
curl -X DELETE http://localhost:8080/api/users/$USER_ID | jq
```

## Desenvolvido para Aprendizado

Este projeto foi desenvolvido como parte de um desafio prático da Rocketseat para iniciantes em Go, focando em:

- Estrutura básica de uma API REST
- Manipulação de HTTP com Go
- Serialização/Desserialização de JSON
- Validação de dados
- Tratamento de erros
- Uso de routers HTTP (Chi)

## Licença

Este projeto é de código aberto e disponível para fins educacionais.
