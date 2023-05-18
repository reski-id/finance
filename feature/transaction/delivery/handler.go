package delivery

import (
	"finance/domain"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	limitUseCase domain.TransactionUseCase
}

func New(nu domain.TransactionUseCase) domain.TransactionHandler {
	return &transactionHandler{
		limitUseCase: nu,
	}
}

func (nh *transactionHandler) InsertTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := nh.limitUseCase.AddTransaction(tmp.ToDomain())
		if err != nil {
			log.Println("Cannot proces data,ID sudah Ada", err)

			c.JSON(http.StatusInternalServerError, "Cannot proces data,ID sudah Ada")
			return (err)

		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    data,
		})
	}
}

func (nh *transactionHandler) GetTransactionID() echo.HandlerFunc {
	return func(c echo.Context) error {

		idTransaction := c.Param("id")
		id, _ := strconv.Atoi(idTransaction)
		data, err := nh.limitUseCase.GetSpecificTransaction(id)

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success data",
			"data":    data,
		})
	}
}
