package delivery

import (
	"finance/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteLimit(e *echo.Echo, bc domain.LimitHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/limit", bc.InsertLimit())
	e.PUT("/limit/:id", bc.UpdateLimit())
	e.DELETE("/limit/:id", bc.DeleteLimit())
	e.GET("/limit/:id", bc.GetLimitID())
}
