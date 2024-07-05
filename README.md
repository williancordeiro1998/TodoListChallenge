# To-Do List Backend

## Introdução

Este projeto é uma implementação do backend de uma lista de tarefas (To-Do List) usando Golang, conforme solicitado pela MoonLabs Studios Inc. O backend utiliza um banco de dados em memória e expõe microserviços utilizando a ferramenta OpenAPI.

## Funcionalidades

- Visualizar lista de tarefas pendentes.
- Visualizar lista de tarefas concluídas.
- Adicionar uma nova tarefa.
- Atualizar o status de uma tarefa (de pendente para concluída e vice-versa, com um limite de 3 vezes para revertê-la a pendente).
- Excluir uma tarefa.

## Regras de Negócio

1. O front-end deve ter acesso às listas de tarefas pendentes e concluídas.
2. Para adicionar uma nova tarefa, deve-se fornecer a descrição, o nome do responsável e o e-mail.
3. Tarefas podem ser movidas de pendentes para concluídas e vice-versa, mas tarefas já concluídas podem ser revertidas para pendentes no máximo 3 vezes.
4. A exclusão de tarefas deve ser física.

## Requisitos Não Funcionais

- Login não é necessário para usar o To-Do.
- Alguma forma de proteção foi usada para proteger os endpoints de acesso público.
- Os métodos HTTP correspondentes a cada operação foram usados.

### Passos para Execução

**Clone o repositório**

git clone <URL do repositório>

cd TodoListChallenge

**Inicialize o módulo Go**

go mod tidy

**Gere a documentação OpenAPI**

swag init

**Execute a aplicação**

go run main.go

Acesse a documentação Swagger em http://localhost:8080/swagger/index.html.
