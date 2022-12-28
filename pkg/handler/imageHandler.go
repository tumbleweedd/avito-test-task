package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	advertisement "github.com/tumbleweedd/avito-test-task"
	"net/http"
	"strconv"
)

func (h *Handler) addImageToAdvertisement(c echo.Context) error {
	advertisementId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "invalid param")
	}

	var img advertisement.Img
	if err := c.Bind(&img); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Println(img.Img)

	imageId, err := h.service.AddImage(advertisementId, img.Img)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"advertisementId": advertisementId,
		"imageId":         imageId,
	})

	return nil
}

type ResponseImagesSet struct {
	Data []string
}

func (h *Handler) getImagesByAdvertisementId(c echo.Context) error {
	advId, err := strconv.Atoi(c.Param("id"))
	var resultImageSet []string

	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "invalid param")
	}

	resultImageSet, err = h.service.Image.GetAllImagesByAdvId(advId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, ResponseImagesSet{
		Data: resultImageSet,
	})

	return nil

}
