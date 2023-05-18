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

	td "finance/feature/transaction/data"
	transactionDelivery "finance/feature/transaction/delivery"
	tu "finance/feature/transaction/usecase"
)

func Initfactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.New(userData, validator)
	userHandler := userDelivery.New(useCase)
	userDelivery.RouteUser(e, userHandler)

	costumerData := dd.New(db)
	contentCase := du.New(costumerData)
	contentHandler := costumerDelivery.New(contentCase)
	costumerDelivery.RouteCostumer(e, contentHandler)

	limitData := ld.New(db)
	limitDCase := lu.New(limitData)
	limitDHandler := limitDelivery.New(limitDCase)
	limitDelivery.RouteLimit(e, limitDHandler)

	transactionData := td.New(db)
	transactionCase := tu.New(transactionData)
	transactionHandler := transactionDelivery.New(transactionCase)
	transactionDelivery.RouteTransaction(e, transactionHandler)
}
