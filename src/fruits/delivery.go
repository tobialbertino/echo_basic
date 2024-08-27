package fruits

import (
	"echo_basic/infra/fruitApi"
	"echo_basic/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Delivery struct {
	Service IService
}

func NewDelivery(service IService) *Delivery {
	return &Delivery{Service: service}
}

func (d *Delivery) GetAllData(c echo.Context) (err error) {
	var res = utils.Response{}

	data, err := d.Service.GetAllData(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	res.Message = utils.Success
	res.Data = data

	return c.JSON(http.StatusOK, res)
}

func (d *Delivery) GetDataByID(c echo.Context) (err error) {
	var res = utils.Response{}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// log
		err = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return
	}

	data, err := d.Service.GetDataByID(c.Request().Context(), id)
	if err != nil {
		// log
		err = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return
	}

	res.Message = utils.Success
	res.Data = data
	return c.JSON(http.StatusOK, res)
}

func (d *Delivery) InsertData(c echo.Context) (err error) {
	var res = utils.Response{}
	newData := fruitApi.DataFruit{}
	err = c.Bind(&newData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = d.Service.InsertData(c.Request().Context(), newData)
	if err != nil {
		// log
		res.Message = utils.MsgFailedToAddWord
		err = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return
	}

	res.Message = utils.WordAdded
	return c.JSON(http.StatusOK, res)
}

func (d *Delivery) DeleteDataByID(c echo.Context) (err error) {
	var res = utils.Response{}
	// get id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// log
		err = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return
	}

	err = d.Service.DeleteDataByID(c.Request().Context(), id)
	if err != nil {
		// log
		err = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return
	}

	res.Message = utils.Success
	return c.JSON(http.StatusOK, res)
}
