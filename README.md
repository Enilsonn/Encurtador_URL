# Encurtador de URL

Este é um projeto simples de encurtador de URLs desenvolvido em Go. Ele expõe uma API REST que recebe uma URL longa e retorna um código encurtado que redireciona para a URL original.

## 📌 Objetivo

O objetivo deste projeto é demonstrar como criar uma API REST com Go usando o pacote `net/http` e o roteador `chi`. O projeto permite:

* Encurtar URLs via POST;
* Redirecionar para URLs originais via GET;
* Manter os dados em memória (sem banco de dados);
* Aplicar boas práticas como uso de middlewares, logs e organização modular de código.

## 🚀 Como rodar o projeto

### 1. Execute o servidor

No terminal, dentro do diretório do projeto, rode:

```bash
go run .
```

Isso iniciará o servidor em `http://localhost:8080`.

### 2. Encurte uma URL com `curl`

Abra um **segundo terminal** e execute o seguinte comando, substituindo `[SUA URL AQUI]` pela URL que deseja encurtar:

```bash
curl -X POST http://localhost:8080/api/shorten -d "{\"url\": \"[SUA URL AQUI]\"}"
```

#### Exemplo:

```bash
curl -X POST http://localhost:8080/api/shorten -d "{\"url\": \"https://google.com\"}"
```

A resposta será algo assim:

```json
{
  "data": "abc123xy"
}
```

### 3. Acesse a URL encurtada no navegador

Copie o valor retornado em `data` e acesse no seu navegador:

```
http://localhost:8080/abc123xy
```

Você será redirecionado para a URL original que você encurtou.

---

## 🧠 Estrutura do Código

* `main.go`: ponto de entrada da aplicação. Inicializa o servidor HTTP e injeta o manipulador de rotas.
* `api.go`: define as rotas da API, os handlers do POST (encurtar) e GET (redirecionar).
* `creteRandomKey.go`: gera um código aleatório de 8 caracteres para identificar cada URL.
* `sendJson.go`: abstração para envio de respostas JSON com tratamento de erro.
* `setApplicationJson.go`: middleware que define o header `Content-Type: application/json` nas respostas.

---

## 🔍 Exemplo de fluxo

1. O usuário envia uma URL via `POST /api/shorten`;
2. O sistema gera um código aleatório e armazena a URL em um mapa em memória;
3. O código é retornado no campo `data`;
4. Quando o usuário acessa `GET /{codigo}`, ele é redirecionado para a URL correspondente.

---

## ℹ️ Observação

> Este projeto **não utiliza banco de dados** — os dados ficam armazenados em um mapa na memória.
> Por isso, todas as URLs encurtadas são perdidas ao reiniciar o servidor.

---

## 📌 Tecnologias usadas

* [Go (Golang)](https://golang.org/)
* [Chi Router](https://github.com/go-chi/chi)
* `net/http` da biblioteca padrão
* JSON como protocolo de comunicação
* Middlewares para logging, recovery e headers

---

Feito com 💻 por Enilson Lima
