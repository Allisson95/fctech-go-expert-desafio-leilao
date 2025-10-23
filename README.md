# Desafio Leilão - Full Cycle Go Expert

Este projeto é uma API de leilão desenvolvida em Go, utilizando MongoDB como banco de dados. A aplicação está pronta para execução via Docker Compose.

## Como executar o projeto

1. **Pré-requisitos:**
   - Docker
   - Docker Compose

2. **Suba os containers:**

```sh
docker-compose up --build
```

A aplicação estará disponível em `http://localhost:8080`.

## Endpoints disponíveis

### Leilão

- **Listar todos os leilões**
  - `GET /auction`
  - Exemplo:
    ```sh
    curl -X GET http://localhost:8080/auction
    ```

- **Buscar leilão por ID**
  - `GET /auction/:auctionId`
  - Exemplo:
    ```sh
    curl -X GET http://localhost:8080/auction/ID_DO_LEILAO
    ```

- **Criar leilão**
  - `POST /auction`
  - Body JSON:
    ```json
    {
      "product_name": "Notebook Dell",
      "category": "Eletrônicos",
      "description": "Notebook i7, 16GB RAM, SSD 512GB",
      "condition": 1
    }
    ```
  - Exemplo:
    ```sh
    curl -X POST http://localhost:8080/auction \
      -H "Content-Type: application/json" \
      -d '{
        "product_name": "Notebook Dell",
        "category": "Eletrônicos",
        "description": "Notebook i7, 16GB RAM, SSD 512GB",
        "condition": 1
      }'
    ```

- **Buscar lance vencedor de um leilão**
  - `GET /auction/winner/:auctionId`
  - Exemplo:
    ```sh
    curl -X GET http://localhost:8080/auction/winner/ID_DO_LEILAO
    ```

### Lance

- **Criar lance**
  - `POST /bid`
  - Body JSON:
    ```json
    {
      "user_id": "ID_DO_USUARIO",
      "auction_id": "ID_DO_LEILAO",
      "amount": 1500.00
    }
    ```
  - Exemplo:
    ```sh
    curl -X POST http://localhost:8080/bid \
      -H "Content-Type: application/json" \
      -d '{
        "user_id": "ID_DO_USUARIO",
        "auction_id": "ID_DO_LEILAO",
        "amount": 1500.00
      }'
    ```

- **Listar lances de um leilão**
  - `GET /bid/:auctionId`
  - Exemplo:
    ```sh
    curl -X GET http://localhost:8080/bid/ID_DO_LEILAO
    ```

### Usuário

- **Buscar usuário por ID**
  - `GET /user/:userId`
  - Exemplo:
    ```sh
    curl -X GET http://localhost:8080/user/ID_DO_USUARIO
    ```

## Observações
- Os parâmetros `ID_DO_LEILAO` e `ID_DO_USUARIO` devem ser substituídos pelos respectivos IDs válidos.
- O MongoDB é inicializado com usuário e senha padrão (`admin`/`admin`).
- As variáveis de ambiente podem ser ajustadas no `docker-compose.yml` conforme necessidade.

---

Projeto desenvolvido para o desafio do curso Go Expert - Full Cycle.
