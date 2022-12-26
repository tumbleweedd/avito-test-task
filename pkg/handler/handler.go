package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/tumbleweedd/avito-test-task/pkg/service"
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
			//advertisement.PUT("/:id", h.updateAdvertisement)
			//advertisement.DELETE("/:id", h.deleteAdvertisement)

			img := advertisement.Group("/:id/image")
			{
				img.POST("/", h.addImage)
				//img.GET("/", h.getImages)
				//img.GET("/:id", h.getImageById)
				//img.PUT("/:id", h.updateImage)
				//img.DELETE("/:id", h.deleteImage)
			}
		}

	}

}