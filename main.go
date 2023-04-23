package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool  `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Learn Go1", Completed: false},
	{ID: "2", Item: "Learn Go2", Completed: false},
	{ID: "3", Item: "Learn Go3", Completed: false},
}

func getTodos (context *gin.Context) {
	context.JSON(http.StatusOK, todos)

}

func addTodos (c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	c.JSON(http.StatusOK, newTodo)
}

func toggleTodoStatus (c *gin.Context) {
	id := c.Param("id")
	t, err := getTodoById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	t.Completed = !t.Completed
	c.JSON(http.StatusOK, t)
}

func getTodoById (id string) (*todo, error){
	for i, t :=range todos {
		if t.ID == id {
			return &todos[i], nil
		}
		
	}
	return nil , errors.New("Todo not found")
}

func getTodo (c *gin.Context){
	id := c.Param("id")
	t, err := getTodoById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, t)
}
func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodos)
	router.Run( "localhost:8080")
}