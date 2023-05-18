package delivery

import (
	"finance/domain"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type costumerHandler struct {
	costumerUseCase domain.CostumerUseCase
}

func New(nu domain.CostumerUseCase) domain.CostumerHandler {
	return &costumerHandler{
		costumerUseCase: nu,
	}
}

func (nh *costumerHandler) InsertCostumer() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := nh.costumerUseCase.AddCostumer(tmp.ToDomain())
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

func (nh *costumerHandler) UpdateCostumer() echo.HandlerFunc {
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

		data, err := nh.costumerUseCase.UpCostumer(cnv, tmp.ToDomain())

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

func (nh *costumerHandler) DeleteCostumer() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		_, errResult := nh.costumerUseCase.DelCostumer(cnv)
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

func (nh *costumerHandler) GetAllCostumer() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := nh.costumerUseCase.GetAllData()

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "Cannot get data no data found")

		}

		if data == nil {
			log.Println("Terdapat error saat mengambil data")
			return c.JSON(http.StatusInternalServerError, "Problem from database")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all data",
			"data":    data,
		})
	}
}

func (nh *costumerHandler) GetCostumerID() echo.HandlerFunc {
	return func(c echo.Context) error {

		idCostumer := c.Param("id")
		id, _ := strconv.Atoi(idCostumer)
		data, err := nh.costumerUseCase.GetSpecificCostumer(id)

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
