# Desafio: API REST com Go, Goroutines e Concorrência

## Objetivo

Criar uma API REST utilizando Go que gerencie um sistema simples de tarefas (To-Do List). O desafio foca no uso de goroutines e concorrência para processar tarefas em segundo plano.

## Requisitos

- [x] **Criar uma API REST** com as seguintes rotas:

  - [x] `POST /tasks` - Criar uma nova tarefa.
  - [x] `GET /tasks` - Listar todas as tarefas.
  - [x] `GET /tasks/{id}` - Obter uma tarefa por ID.
  - [x] `DELETE /tasks/{id}` - Excluir uma tarefa por ID.

- [x] **Modelo de Tarefa**

  - [x] `id` (UUID ou inteiro auto-incremental)
  - [x] `title` (string)
  - [x] `description` (string)
  - [x] `completed` (boolean)
  - [x] `created_at` (timestamp)

- [ ] **Concorrência**
  - [x] Implementar um sistema de processamento de tarefas em segundo plano utilizando goroutines e channels.
  - [x] Criar uma rota `POST /tasks/{id}/process` que inicia o processamento da tarefa em uma goroutine, simulando um trabalho assíncrono (exemplo: um `time.Sleep(3 * time.Second)` antes de marcar a tarefa como concluída).

## Requisitos Técnicos

- [x] Utilizar `net/http` ou `gin-gonic/gin` para criar a API.
- [ ] Utilizar `sync.Mutex` ou outro mecanismo de concorrência para evitar condições de corrida ao acessar as tarefas.
- [ ] Armazenar as tarefas em uma estrutura de dados em memória (mapa ou slice com mutex).
- [x] Implementar logs para acompanhar o processamento das tarefas.

## Diferenciais (Não obrigatórios, mas valorizados)

- [ ] Implementação de testes unitários para os handlers da API.
- [ ] Uso de um banco de dados (SQLite, PostgreSQL ou MongoDB) em vez de armazenamento em memória.
- [ ] Uso de `context.Context` para controlar o tempo de execução das goroutines.

## Entrega

- [x] Criar um repositório público no GitHub contendo o código-fonte e instruções no README sobre como rodar a aplicação.
- [ ] Opcional: Criar um `Dockerfile` para facilitar a execução do projeto.
