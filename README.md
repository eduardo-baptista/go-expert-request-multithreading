# Request Multithreading

Desafio realizado para o curso [Go-Expert](https://goexpert.fullcycle.com.br/pos-goexpert/)

## Descrição

O desafio consiste em criar uma aplicação que faça uma requisição HTTP para as URLs:
- https://viacep.com.br
- https://cdn.apicep.com

retornando o resultado da que responda mais rápido.

Caso o tempo de resposta seja maior que 1 segundo, a aplicação deve exibir o erro de timeout.

## Como executar

```bash
go run main.go 00000-000
```