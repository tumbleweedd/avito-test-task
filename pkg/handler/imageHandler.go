package handler

import (
	"fmt"
	advertisement "github.com/tumbleweedd/avito-test-task/model"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
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

func (h *Handler) getImageById(c echo.Context) error {
	var response advertisement.ImageResponse

	advId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid param")
	}

	imgId, err := strconv.Atoi(c.Param("imgId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid param")
	}

	response, err = h.service.Image.GetImageById(advId, imgId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, advertisement.ImageResponse{
		Image: response.Image,
	})

	return nil
}

func (h *Handler) deleteImage(c echo.Context) error {
	advId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid advertisement param")
	}

	imageId, err := strconv.Atoi(c.Param("imgId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid image param")
	}

	err = h.service.Image.DeleteImage(advId, imageId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "the deletion was successful",
	})

	return nil
}
