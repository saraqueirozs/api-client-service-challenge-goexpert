# 1 - Desafio: Client-Server-API

## Pré-requisitos

- [Go](https://golang.org/doc/install) (versão 1.22 ou superior)
- [SQLite](https://www.sqlite.org/download.html)

## Passo a Passo

Após o clone do repositório, siga os próximos passos!

1. Possível problema GCC windows, siga essa possível solução:
Faça o download para instalação do GCC MinGW:
MinGW 64 bits: https://sourceforge.net/projects/ming...
MinGW 64 bits(4.8.5): https://sourceforge.net/projects/ming...
MinGW 32 bits: https://sourceforge.net/projects/ming...

Inclua no PATH das variáveis de ambiente do seu computador. Após estes ajustes, faça um split no seu terminal. Um para rodar o server.go e outro para rodar o client.go 

    - O banco de dados será criado automaticamente ao rodar o `server/main.go`.
    - O servidor estará na porta `8080`.

  ```sh
   cd server
   go run server.go
   ```

  ```sh
   cd client
   go run client.go
   ```

2. Em outro terminal, execute o cliente:
   ```sh
   cd client
   go run main.go
   ```
    - O cliente fará uma requisição ao servidor para obter a cotação do dólar e salvará o valor no arquivo `cotacao.txt`.