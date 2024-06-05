# RamenGO! RedVentures


## Estrutura do projeto
Contém um banco de dados MySQL e uma aplicação escrita
em Golang que se comunica com o mundo exterior mediante
seus endpoints.

## How to...
Esta sessão contém instruções para rodar o projeto
localmente usando docker-compose. 
### Run the project
#### Pré-requisitos
- Docker
- Docker-compose

As únicas dependências para rodar o projeto são ter em seu
ambiente o Docker e o Docker-compose instalados.

#### Rodando o projeto
Para rodar o projeto, basta executar o comando abaixo estando
na pasta raiz do projeto (/RamenGORV).
```bash
docker-compose up -d
```
**Assista a mágica acontecer** 

## Não conseguiu?

## Endpoints
Os endpoints estão disponíveis na pasta ```/pkg/test/local.http```
e para testá-los, basta utilizar a extensão REST Client do VSCode.
