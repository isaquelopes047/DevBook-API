# 📘 DevBook-API

**DevBook-API** é uma aplicação real desenvolvida em [Golang (Go)](https://golang.org), que representa a estrutura completa de backend para uma **rede social**. A API foi construída com foco em desempenho, segurança, boas práticas REST e escalabilidade, utilizando autenticação JWT, controle de usuários, publicações, relações sociais (seguidores) e muito mais.

---

## 🚀 Visão Geral

A DevBook-API fornece todas as funcionalidades principais necessárias para operar uma rede social moderna, incluindo:

- Autenticação segura com tokens JWT
- Gerenciamento de usuários
- Relacionamentos de seguir/deixar de seguir
- CRUD completo de publicações
- Feed personalizado baseado nos usuários seguidos
- Camadas bem definidas (router, controller, model, middleware)
- Validações e criptografia de senhas

---

## 🛠️ Tecnologias Utilizadas

- **Go (Golang)**
- **Gorilla Mux** — roteador HTTP
- **MySQL** — banco de dados relacional
- **bcrypt** — hash seguro de senhas
- **JWT (JSON Web Tokens)** — autenticação
- **Docker** (opcional) — para containerização e setup de ambiente

---

## 📚 Funcionalidades

### 👥 Usuários
- Cadastro, login, atualização e remoção
- Busca por nome ou nickname
- Seguir e deixar de seguir usuários
- Listagem de seguidores e seguidos

### 🔐 Autenticação
- Login com verificação de senha e geração de JWT
- Middleware para proteção de rotas autenticadas

### 📝 Publicações
- Criar, editar e excluir publicações
- Listar publicações próprias ou de usuários seguidos
- Curtidas e descurtidas (opcional)
