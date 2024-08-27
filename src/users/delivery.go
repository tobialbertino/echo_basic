package users

import (
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
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// log
		err = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return
	}

	data, err := d.Service.GetByID(c.Request().Context(), id)
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
	newData := UserEntity{}
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
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// log
		err = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return
	}
	// get data payload
	newData := UserEntity{}
	err = c.Bind(&newData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = d.Service.UpdateDataByID(c.Request().Context(), id, newData)
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
