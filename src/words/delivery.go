package words

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

func (d *Delivery) GetAllData(c echo.Context) (err error) {
	var res = utils.Response{}

	data := d.Service.GetAllData(c.Request().Context())
	res.Message = utils.AllWords
	res.Data = data

	return c.JSON(http.StatusOK, res)
}

func (d *Delivery) InsertData(c echo.Context) (err error) {
	var res = utils.Response{}
	newData := ReqInsertWord{}
	err = c.Bind(&newData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = d.Service.InsertData(c.Request().Context(), newData.Word)
	if err != nil {
		// log
		res.Message = utils.MsgFailedToAddWord
		err = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return
	}

	res.Message = utils.WordAdded
	return c.JSON(http.StatusOK, res)
}
