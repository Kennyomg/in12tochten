package routes

import (
	"compress/gzip"
	"github.com/kennyomg/in12tochten/server/api/todo/routes"
	
	"github.com/kennyomg/in12tochten/server/common/static"
	
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: gzip.BestCompression,
	})	
	
	static.Init(e)
	
	todoroutes.Init(e)
}
