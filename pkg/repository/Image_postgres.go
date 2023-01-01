package repository

import (
	"fmt"
	advertisement "github.com/tumbleweedd/avito-test-task/model"

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
	addImageQuery := fmt.Sprintf("insert into img (img) values ($1) returning id")

	row := tx.QueryRow(addImageQuery, image)
	err = row.Scan(&imageId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createAdvertisementImageQuery := fmt.Sprintf("insert into advertisement_img (advertisement_id, img_id) values ($1, $2)")
	_, err = tx.Exec(createAdvertisementImageQuery, advertisementId, imageId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return imageId, tx.Commit()
}

func (r *ImagePostgres) GetAllImagesByAdvId(advId int) ([]string, error) {
	var result []string

	query := fmt.Sprint(`
						select img.img
						from img
						where img.id in (select ai.img_id
											from advertisement_img ai
											where ai.img_id in (select ai2.img_id from advertisement_img ai2 where ai2.advertisement_id = $1))
							`)

	if err := r.db.Select(&result, query, advId); err != nil {
		return nil, err
	}
	return result, nil

}

func (r *ImagePostgres) GetImageById(advId, imageId int) (advertisement.ImageResponse, error) {
	var image advertisement.ImageResponse
	query := fmt.Sprintf(`
						SELECT i.img
							FROM %s i
							JOIN %s ai on ai.img_id = i.id
							where ai.advertisement_id = $1 and ai.img_id = $2
						`, imgTable, advertisementImgTable)

	err := r.db.Get(&image.Image, query, advId, imageId)

	return image, err
}

func (r *ImagePostgres) DeleteImage(advId, imageId int) error {
	query := fmt.Sprintf(`delete
							FROM img i
							WHERE i.id in (
								SELECT ai.img_id
									FROM advertisement_img ai
									WHERE ai.advertisement_id = $1 and ai.img_id = $2
							);
						`)

	_, err := r.db.Exec(query, advId, imageId)
	if err != nil {
		return err
	}

	return nil
}
