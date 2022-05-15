package routes

import (
	"github.com/joshua-ather/sv_be/app/controllers"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/article", controllers.CreatePost)
	e.GET("/article/:id", controllers.GetPostById)
	e.PUT("/article/:id", controllers.UpdatePost)
	e.DELETE("/article/:id", controllers.DeletePost)
	e.GET("/article/:limit/:offset", controllers.GetPosts)

	return e
}
