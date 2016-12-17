package todoroutes

import (
	"github.com/kennyomg/in12tochten/server/api/todo/controller"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	e.GET("/api/todos", todocontroller.GetAll)
	e.POST("/api/todos", todocontroller.NewTodo)
	e.DELETE("/api/todos/:id", todocontroller.RemoveTodo)
}
