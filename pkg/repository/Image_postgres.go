package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ImagePostgres struct {
	db *sqlx.DB
}

func NewImagePostgres(db *sqlx.DB) *ImagePostgres {
	return &ImagePostgres{db: db}
}

func (r *ImagePostgres) AddImage(advertisementId int, image string) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var imageId int
	addImageQuery := fmt.Sprintf("insert into %s (img) values ($1) returning id", imgTable)

	row := tx.QueryRow(addImageQuery, image)
	err = row.Scan(&imageId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createAdvertisementImageQuery := fmt.Sprintf("insert into %s (advertisement_id, img_id) values ($1, $2)", advertisementImgTable)
	_, err = tx.Exec(createAdvertisementImageQuery, advertisementId, imageId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return imageId, tx.Commit()
}
