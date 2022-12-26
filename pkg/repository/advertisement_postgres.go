package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/avito-test-task"
)

type AdvertisementPostgres struct {
	db *sqlx.DB
}

func NewAdvertisementPostgres(db *sqlx.DB) *AdvertisementPostgres {
	return &AdvertisementPostgres{
		db: db,
	}
}

func (r *AdvertisementPostgres) CreateAdvertisement(input advertisement.Advertisement) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var advertisementId int
	createAdvertisementQuery := fmt.Sprintf("insert into %s (title, description) values($1, $2) returning id", advertisementTable)

	rowA := tx.QueryRow(createAdvertisementQuery, input.Title, input.Description)
	err = rowA.Scan(&advertisementId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var imgId int
	createImgQuery := fmt.Sprintf("insert into %s (img) values($1) returning id", imgTable)

	rowA = tx.QueryRow(createImgQuery, input.Img)
	err = rowA.Scan(&imgId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createAdvertisementImgQuery := fmt.Sprintf("insert into %s (advertisement_id, img_id) values ($1,$2)", advertisementImgTable)
	_, err = tx.Exec(createAdvertisementImgQuery, advertisementId, imgId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return advertisementId, tx.Commit()
}

func (r *AdvertisementPostgres) GetAllAdvertisement() ([]advertisement.Advertisement, error) {
	var advertisements []advertisement.Advertisement

	query := fmt.Sprintf(`
								SELECT a.id,
									   a.description,
									   a.title,
									   i.img
									FROM %s a
									join %s ai on a.id = ai.advertisement_id
									join %s i on i.id = ai.img_id
									ORDER  BY i.id, a.id
									limit (select count(*) from advertisement)
									`, advertisementTable, advertisementImgTable, imgTable)
	if err := r.db.Select(&advertisements, query); err != nil {
		return nil, err
	}

	return advertisements, nil
}

func (r *AdvertisementPostgres) GetAdvertisementById(id int) (advertisement.AdvertisementDTO, error) {
	var advertisement advertisement.AdvertisementDTO

	query := fmt.Sprintf(`
								SELECT a.description,
									   a.title,
									   i.img
									FROM %s a
									join %s ai on a.id = ai.advertisement_id
									join %s i on i.id = ai.img_id
									where a.id = $1
									ORDER  BY i.id, a.id
									limit (select count(*) from advertisement)
									`, advertisementTable, advertisementImgTable, imgTable)

	err := r.db.Get(&advertisement, query, id)

	return advertisement, err

}
