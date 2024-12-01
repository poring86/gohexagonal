# GoHexagonal

GoHexagonal é uma aplicação de exemplo que implementa a arquitetura hexagonal (Ports and Adapters) usando Go. O objetivo é demonstrar como construir uma aplicação modular e extensível, separando a lógica de negócios das implementações específicas de infraestrutura.

---

## 🚀 Funcionalidades

- Gerenciamento de produtos:
  - Criar produtos.
  - Habilitar ou desabilitar produtos.
  - Buscar produtos por ID.
- API RESTful implementada com Gorilla Mux.
- Middlewares utilizando Negroni.

---

## 🛠️ Tecnologias Utilizadas

- [Go](https://golang.org/)
- [Gorilla Mux](https://github.com/gorilla/mux) para o roteamento.
- [Negroni](https://github.com/urfave/negroni) para middlewares.
- SQLite3 como banco de dados.
- Docker e Docker Compose para ambiente de desenvolvimento.

---

## 📦 Instalação e Configuração

### Pré-requisitos

- [Go](https://golang.org/doc/install) 1.18 ou superior.
- [Docker](https://docs.docker.com/get-docker/) e [Docker Compose](https://docs.docker.com/compose/install/) (opcional, para rodar via container).

### Clone o Repositório

```bash
git clone https://github.com/poring86/gohexagonal.git
cd gohexagonal
