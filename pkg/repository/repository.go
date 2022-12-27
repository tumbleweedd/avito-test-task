package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/avito-test-task"
)

type Advertisement interface {
	GetAllAdvertisement() ([]advertisement.Advertisement, error)
	CreateAdvertisement(input advertisement.Advertisement) (int, error)
	GetAdvertisementById(id int) (advertisement.AdvertisementDTO, error)
	UpdateAdvertisement(id int, dto advertisement.UpdateAdvertisement) error
}

type Image interface {
	AddImage(advertisementId int, image string) (int, error)
}

type Repository struct {
	Advertisement
	Image
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Advertisement: NewAdvertisementPostgres(db),
		Image:         NewImagePostgres(db),
	}
}
