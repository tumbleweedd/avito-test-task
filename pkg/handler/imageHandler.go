package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	advertisement "github.com/tumbleweedd/avito-test-task"
	"net/http"
	"strconv"
)

func (h *Handler) addImage(c echo.Context) error {
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
