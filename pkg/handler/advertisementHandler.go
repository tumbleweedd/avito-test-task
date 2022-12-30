package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	advertisement "github.com/tumbleweedd/avito-test-task/model"
)

func (h *Handler) createAdvertisement(c echo.Context) error {
	var input advertisement.Advertisement
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	err := input.Validate()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

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
	adv, err := h.service.Advertisement.GetAllAdvertisement(getSortValue(c))
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

func (h *Handler) updateAdvertisement(c echo.Context) error {
	advId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var updatedAdv advertisement.UpdateAdvertisement

	if err := c.Bind(&updatedAdv); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.service.UpdateAdvertisement(advId, updatedAdv); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})

	return nil
}

func (h *Handler) deleteAdvertisement(c echo.Context) error {
	advId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.service.DeleteAdvertisement(advId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "the deletion was successful",
	})
	return nil
}

func getSortValue(c echo.Context) string {
	var sortBy string

	val := c.Request().Header.Get("sort-by")

	if strings.EqualFold(val, "desc") {
		sortBy = "desc"
	} else if val == "" {
		sortBy = ""
	}

	return sortBy
}
