package factory

import (
	ud "finance/feature/user/data"
	userDelivery "finance/feature/user/delivery"
	us "finance/feature/user/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	dd "finance/feature/costumer/data"
	costumerDelivery "finance/feature/costumer/delivery"
	du "finance/feature/costumer/usecase"

	ld "finance/feature/limit/data"
	limitDelivery "finance/feature/limit/delivery"
	lu "finance/feature/limit/usecase"
)

func Initfactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.New(userData, validator)
	userHandler := userDelivery.New(useCase)
	userDelivery.RouteUser(e, userHandler)

	costumerData := dd.New(db)
	contentDCase := du.New(costumerData)
	contentDHandler := costumerDelivery.New(contentDCase)
	costumerDelivery.RouteCostumer(e, contentDHandler)

	limitData := ld.New(db)
	limitDCase := lu.New(limitData)
	limitDHandler := limitDelivery.New(limitDCase)
	limitDelivery.RouteLimit(e, limitDHandler)
}
