package delivery

import (
	"finance/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteCostumer(e *echo.Echo, bc domain.CostumerHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/costumer", bc.InsertCostumer())
	e.PUT("/costumer/:id", bc.UpdateCostumer())
	e.DELETE("/costumer/:id", bc.DeleteCostumer())
	e.GET("/costumer", bc.GetAllCostumer())
	e.GET("/costumer/:id", bc.GetCostumerID())
}
