package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	advertisement "github.com/tumbleweedd/avito-test-task"
	"net/http"
	"strconv"
)

func (h *Handler) createAdvertisement(c echo.Context) error {
	var input advertisement.Advertisement
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	fmt.Println(input.Img)

	id, err := h.service.Advertisement.CreateAdvertisement(input)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

	return nil
}

type getAllAdvertisementResponse struct {
	Data []advertisement.Advertisement `json:"data"`
}

func (h *Handler) getAllAdvertisement(c echo.Context) error {
	adv, err := h.service.Advertisement.GetAllAdvertisement()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, getAllAdvertisementResponse{
		Data: adv,
	})
	return nil
}

func (h *Handler) getAdvertisementById(c echo.Context) error {
	advId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	advDTO, err := h.service.Advertisement.GetAdvertisementById(advId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, advDTO)
	return nil
}
