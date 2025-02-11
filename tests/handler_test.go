package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list/internal/data"
	"todo-list/internal/handler"
	"todo-list/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	// Preparar o ambiente de testes
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Simular algumas tarefas para testar
	data.Tasks = []models.Task{
		{ID: 1, Title: "Test Task 1", Completed: false},
		{ID: 2, Title: "Test Task 2", Completed: true},
	}

	// Registrar o endpoint
	router.GET("/tasks", handler.GetTasks)

	// Criar uma requisição de teste
	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()

	// Processar a requisição
	router.ServeHTTP(w, req)

	// Verificar o status de resposta e conteúdo
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Task 1")
	assert.Contains(t, w.Body.String(), "Test Task 2")
}