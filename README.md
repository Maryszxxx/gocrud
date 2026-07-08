# GoCRUD

API REST em Go para gerenciamento de usuários, construída com [Gin](https://github.com/gin-gonic/gin), [MongoDB](https://www.mongodb.com/) (driver v2) e arquitetura em camadas (Controller → Service → Repository).

Projeto desenvolvido para estudo e prática de backend em Go, cobrindo desde CRUD básico até autenticação simples de usuários.

## funcionalidades

- Cadastro de usuários com validação de dados (nome, email, senha, idade)
- Autenticação (login) por email e senha
- Busca de usuário por ID ou por email
- Atualização parcial de dados do usuário (nome e idade)
- Remoção de usuário
- Criptografia de senha (MD5)
- Logs estruturados com [Zap](https://github.com/uber-go/zap)
- Tratamento de erros padronizado (`rest_err`)
- Validação de payloads com mensagens traduzidas (`go-playground/validator`)

## arquitetura

O projeto segue uma separação em camadas:

```
main.go
├── controller/        # Recebe requisições HTTP, valida entrada e formata resposta
│   └── routes/         # Definição das rotas
├── model/
│   ├── service/         # Regras de negócio
│   └── repository/      # Acesso ao banco de dados (MongoDB)
│       └── entity/       # Entidades persistidas + conversores
├── config/
│   ├── database/mongodb/ # Conexão com o MongoDB
│   ├── logger/            # Configuração de logs (Zap)
│   ├── rest_err/          # Erros HTTP padronizados
│   └── validation/        # Tradução de erros de validação
└── view/                # Conversão de domínio para resposta HTTP
```

Cada camada depende apenas de interfaces da camada abaixo, o que facilita testes e manutenção.

## tecnologias

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin) — framework HTTP
- [MongoDB Go Driver v2](https://www.mongodb.com/docs/drivers/go/current/)
- [Zap](https://github.com/uber-go/zap) — logging estruturado
- [Validator v10](https://github.com/go-playground/validator) — validação de structs
- [godotenv](https://github.com/joho/godotenv) — variáveis de ambiente

## ⚙️ Pré-requisitos

- Go instalado (versão compatível conforme `go.mod`)
- MongoDB rodando localmente ou via Docker:

```bash
docker run -d -p 27017:27017 --name mongo-gocrud mongo
```

## 🔧 Configuração

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```env
MONGODB_URL=mongodb://localhost:27017
MONGODB_USER_DB=users
MONGODB_USER_COLLECTION=users
```

## ▶️ Executando o projeto

```bash
go mod tidy
go run .
```

> ⚠️ Importante: use `go run .` (rodando o pacote inteiro) e não `go run main.go`, pois o projeto possui múltiplos arquivos no `package main` (ex: `init_dependencies.go`).

O servidor sobe por padrão em `http://localhost:8080`.

## 📚 Endpoints

| Método | Rota                          | Descrição                       |
|--------|-------------------------------|----------------------------------|
| POST   | `/createUser`                 | Cria um novo usuário             |
| POST   | `/login`                      | Autentica um usuário             |
| GET    | `/getUserById/:userId`        | Busca usuário por ID             |
| GET    | `/getUserByEmail/:email`      | Busca usuário por email          |
| PUT    | `/updateUser/:userId`         | Atualiza nome e/ou idade         |
| DELETE | `/deleteUser/:userId`         | Remove um usuário                |

### Exemplo — Criar usuário

```http
POST /createUser
Content-Type: application/json

{
  "name": "Maria",
  "email": "maria@exemplo.com",
  "password": "Senha123",
  "age": 25
}
```

### Exemplo — Login

```http
POST /login
Content-Type: application/json

{
  "email": "maria@exemplo.com",
  "password": "Senha123"
}
```

### Exemplo — Atualizar usuário

```http
PUT /updateUser/{userId}
Content-Type: application/json

{
  "name": "Mario Jorge",
  "age": 26
}
```

## estrutura de resposta (usuário)

```json
{
  "id": "64f1a2b3c4d5e6f7a8b9c0d1",
  "name": "Maria",
  "email": "maria@exemplo.com",
  "age": 25
}
```

## testando a API

Você pode testar os endpoints usando [Postman](https://www.postman.com/), [Insomnia](https://insomnia.rest/) ou `curl`.

## 📌 Roadmap / Melhorias futuras

- [ ] Autenticação com JWT
- [ ] Testes automatizados (unitários e de integração)
- [ ] Migrar hash de senha de MD5 para bcrypt
- [ ] Documentação via Swagger/OpenAPI
- [ ] Containerização com Docker Compose (API + MongoDB)
