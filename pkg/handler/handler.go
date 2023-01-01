package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/tumbleweedd/avito-test-task/pkg/service"
)

const (
	ascendingDateSort   = "date-asc"
	descendingDateSort  = "date-desc"
	ascendingPriceSort  = "price-asc"
	descendingPriceSort = "price-desc"
	getAllImages        = "all"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes(e *echo.Echo) {
	api := e.Group("/api")
	{
		advertisement := api.Group("/advertisement")
		{
			advertisement.POST("/", h.createAdvertisement)
			advertisement.GET("/", h.getAllAdvertisement)
			advertisement.GET("/:id", h.getAdvertisementById)
			advertisement.PUT("/:id", h.updateAdvertisement)
			advertisement.DELETE("/:id", h.deleteAdvertisement)

			img := advertisement.Group("/:id/image")
			{
				img.POST("/", h.addImageToAdvertisement)
				img.GET("/", h.getImagesByAdvertisementId)
				img.GET("/:imgId", h.getImageById)
				// img.PUT("/:id", h.updateImage)
				img.DELETE("/:imgId", h.deleteImage)
			}
		}

	}

}
