package infrastructure

import (
	controllers "echoSample/src/interfaces/api"

	"github.com/labstack/echo"
)

func Init() {
	// Echo instance
	e := echo.New()
	userController := controllers.NewUserController(NewSqlHandler())

	e.GET("/usres", func(c echo.Context) error)
	useres := userCo
}
