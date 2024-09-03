package products

import (
	"echo_basic/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Delivery struct {
	Service IService
}

func NewDelivery(service IService) *Delivery {
	return &Delivery{Service: service}
}

func (d *Delivery) GetListData(c echo.Context) (err error) {
	var res = utils.Response{}

	data := d.Service.GetListData(c.Request().Context())
	res.Message = utils.Success
	res.Data = data

	return c.JSON(http.StatusOK, res)
}

func (d *Delivery) GetByID(c echo.Context) (err error) {
	var res = utils.Response{}
	idStr := c.Param("id")

	data, err := d.Service.GetByID(c.Request().Context(), idStr)
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
	newData := ProductEntity{}
	err = c.Bind(&newData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = d.Service.InsertData(c.Request().Context(), newData)
	if err != nil {
		// log
		err = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return
	}

	res.Message = utils.Success
	return c.JSON(http.StatusOK, res)
}

func (d *Delivery) UpdateDataByID(c echo.Context) (err error) {
	var res = utils.Response{}
	// get id
	idStr := c.Param("id")

	// get data payload
	newData := ProductEntity{}
	err = c.Bind(&newData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newData.ID = idStr
	err = d.Service.UpdateDataByID(c.Request().Context(), idStr, newData)
	if err != nil {
		// log
		err = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return
	}

	res.Message = utils.Success
	return c.JSON(http.StatusOK, res)
}

func (d *Delivery) DeleteDataByID(c echo.Context) (err error) {
	var res = utils.Response{}
	// get id
	idStr := c.Param("id")

	err = d.Service.DeleteDataByID(c.Request().Context(), idStr)
	if err != nil {
		// log
		err = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return
	}

	res.Message = utils.Success
	return c.JSON(http.StatusOK, res)
}
