package delivery

import (
	"finance/domain"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type limitHandler struct {
	limitUseCase domain.LimitUseCase
}

func New(nu domain.LimitUseCase) domain.LimitHandler {
	return &limitHandler{
		limitUseCase: nu,
	}
}

func (nh *limitHandler) InsertLimit() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := nh.limitUseCase.AddLimit(tmp.ToDomain())
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

func (nh *limitHandler) UpdateLimit() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		var tmp InsertRequest
		res := c.Bind(&tmp)

		if res != nil {
			log.Println(res, "Cannot parse data")
			return c.JSON(http.StatusInternalServerError, "error read update")
		}

		data, err := nh.limitUseCase.UpLimit(cnv, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot update data", err)
			c.JSON(http.StatusInternalServerError, "cannot update")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update data",
			"data":    data,
		})
	}
}

func (nh *limitHandler) DeleteLimit() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		_, errResult := nh.limitUseCase.DelLimit(cnv)
		if errResult != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  errResult.Error(),
				"message": "Cannot deleted data",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete data",
		})
	}
}

func (nh *limitHandler) GetLimitID() echo.HandlerFunc {
	return func(c echo.Context) error {

		idLimit := c.Param("id")
		id, _ := strconv.Atoi(idLimit)
		data, err := nh.limitUseCase.GetSpecificLimit(id)

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
