package handler

import (
	"fmt"
	"net/http"
	"strconv"
	// "strings"

	"github.com/labstack/echo/v4"
	advertisement "github.com/tumbleweedd/avito-test-task/model"
	"strings"
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
	limitParam, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid limit param")
	}

	fmt.Println(limitParam)

	offsetParam, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid offset param")
	}

	sortBy, err := getSortValue(c)
	if err != nil {
		return err
	}

	adv, err := h.service.Advertisement.GetAllAdvertisement(sortBy, limitParam, offsetParam*5)
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

	advWithAllImg := new(advertisement.AdvertisemenWithAllImgtDTO)
	advWithAllImg.Title = advDTO.Title
	advWithAllImg.Description = advDTO.Description
	advWithAllImg.Img, _ = h.service.Image.GetAllImagesByAdvId(advId)
	advWithAllImg.DateTime = advDTO.DateTime
	advWithAllImg.Price = advDTO.Price

	return advertisementResponse(c, advWithAllImg, &advDTO)
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

func getSortValue(c echo.Context) (string, error) {
	var sortBy string

	val := c.Request().Header.Get("sort-by")

	if strings.EqualFold(val, descendingDateSort) {
		sortBy = "v.date_creation desc"
	} else if strings.EqualFold(val, ascendingDateSort) {
		sortBy = "v.date_creation"
	} else if strings.EqualFold(val, descendingPriceSort) {
		sortBy = "v.price desc"
	} else if strings.EqualFold(val, ascendingPriceSort) {
		sortBy = "v.price"
	} else if val == "" {
		sortBy = "v.date_creation desc"
	} else {
		return "", echo.NewHTTPError(http.StatusBadRequest, "invalid header")
	}

	return sortBy, nil
}

func advertisementResponse(
	c echo.Context,
	advAllImgDTO *advertisement.AdvertisemenWithAllImgtDTO,
	advDTO *advertisement.AdvertisementDTO) error {
	val := c.Request().Header.Get("fields")

	if strings.EqualFold(val, getAllImages) {
		c.JSON(http.StatusOK, advAllImgDTO)
		return nil
	} else if val == "" {
		c.JSON(http.StatusOK, *advDTO)
		return nil
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid header")
	}
}

func paginationAdv(c echo.Context) {
	c.Request().Header.Add("limit", "5")
}
