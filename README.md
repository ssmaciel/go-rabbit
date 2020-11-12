<h1 align="center">
    🚀 Go RabbitMQ
</h1>

<p align="center">Microserviços com RabbitMQ</p>

<p align="center">
  <img src="https://img.shields.io/static/v1?label=golang&message=1.15&color=7fd5ea&logo=go" />
  <img src="https://img.shields.io/static/v1?label=golang&message=1.15&color=FF6600&logo=rabbitmq" />
  <img src="https://img.shields.io/static/v1?label=docker&message=19.03.13&color=0073ec&logo=docker" />
  <img src="https://img.shields.io/static/v1?label=kubernets&message=19.03.13&color=326CE5&logo=kubernetes" />
  <img src="https://img.shields.io/badge/last%20commit-november-important" />
  <img src="https://img.shields.io/badge/license-MIT-success"/>
</p>

<p align="center">
  <a href="#-features">Features</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-pré-requisitos">Pré-Requisitos</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-tecnologias">Tecnologias</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-licença">Lincença</a>
</p>

<h3 align="center"> 
🚧  Finalizado  🚧
</h3>

### 📎 Features 

- [x] Criação de Microserviços
- [x] Comunicação entre microserviços
- [x] Utilização de Filas com RabbitMQ
- [x] Criação de Containers e Images com Docker
- [x] Gerenciamento de Containers com Kurbenetes

### 💻 Projeto

Desenvolvido duas aplicações com Microsserviços utilizando Golang, Containers com Docker e Kurbenetes

### ✅ Demonstração do Frontend
<img src="https://github.com/ssmaciel/go-rabbit/blob/main/assets/front-end.png" />

### ⚙ Pré-requisitos

Antes de começar, você vai precisar ter instalado em sua máquina as seguintes ferramentas:
[Git](https://git-scm.com), [Docker](https://www.docker.com/), [Golang](https://golang.org/) (caso queira debugar o fonte é claro rs).
Além disto é bom ter um editor para trabalhar com o código como [VSCode](https://code.visualstudio.com/)


### 📙 Rodando os Microsserviços (Ambiente completo)

```bash
# Clone este repositório
$ git clone https://github.com/ssmaciel/go-rabbit.git

# Rode o docker-compose para subir o ambiente
$ docker-compose up -d

```

### 📗 Rodando os Microsserviços (Ambientes separados)

```bash
# Clone este repositório
$ git clone https://github.com/ssmaciel/go-rabbit.git

# Navegue para a pasta do rabbit
$ cd rabbit

# Rode o docker-compose para subir rabbit
$ docker-compose up -d

# Navegue para a pasta do consumer
$ cd ../consumer

# Rode o microserviço consumer
$ go run consumer.go

# Navegue para a pasta do producer
$ cd ../producer

# Rode o microserviço consumer
$ go run producer.go
```

### 🚀 Tecnologias

Esse projeto foi desenvolvido com as seguintes tecnologias:

- Golang
- RabbitMQ

### 📕 Bibliotecas

Esse projeto foi utilizou das seguintes lib:

- godotenv
- go-retryablehttp
- go-rabbitmq

### 📙 Arquitetura do Projeto

Para uma melhorar estrutura de projetos utilizamos das seguintes fundamentos:

- DDD
- CI & CD

###  📘 Padrão de Código

Para padronizar a escrita do código, utilizamos as seguinte ferramentas:

- Eslint
- Prettier
- EditorConfig


### 📝 Licença

Esse projeto está sob a licença MIT.

<hr/>

Feito por Samuel Maciel
