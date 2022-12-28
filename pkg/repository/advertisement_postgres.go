package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/avito-test-task"
)

type AdvertisementPostgres struct {
	db *sqlx.DB
}

func (r *AdvertisementPostgres) DeleteAdvertisement(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	queryForAdv := fmt.Sprint(`
									delete from advertisement a 
       								where a.id = $1
       								`)

	queryForImg := fmt.Sprint(`
									delete
									from img
									where img.id in (select ai.img_id
													 from advertisement_img ai
													 where ai.img_id in (select ai2.img_id from advertisement_img ai2 where ai2.advertisement_id = $1));
									`)

	_, err = tx.Exec(queryForImg, id)
	if err != nil {
		return err
		tx.Rollback()
	}

	_, err = tx.Exec(queryForAdv, id)
	if err != nil {
		return err
		tx.Rollback()
	}

	return tx.Commit()
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
	createAdvertisementQuery := fmt.Sprint("insert into advertisement (title, description) values($1, $2) returning id")

	rowA := tx.QueryRow(createAdvertisementQuery, input.Title, input.Description)
	err = rowA.Scan(&advertisementId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var imgId int
	createImgQuery := fmt.Sprint("insert into img (img) values($1) returning id")

	rowA = tx.QueryRow(createImgQuery, input.Img)
	err = rowA.Scan(&imgId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createAdvertisementImgQuery := fmt.Sprint("insert into advertisement_img (advertisement_id, img_id) values ($1,$2)")
	_, err = tx.Exec(createAdvertisementImgQuery, advertisementId, imgId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return advertisementId, tx.Commit()
}

func (r *AdvertisementPostgres) GetAllAdvertisement() ([]advertisement.Advertisement, error) {
	var advertisements []advertisement.Advertisement

	query := fmt.Sprint(`
								SELECT a.id,
									   a.description,
									   a.title,
									   i.img
									FROM advertisement a
									join advertisement_img ai on a.id = ai.advertisement_id
									join img i on i.id = ai.img_id
									ORDER  BY a.id desc
									limit (select count(*) from advertisement)
									`)
	if err := r.db.Select(&advertisements, query); err != nil {
		return nil, err
	}

	return advertisements, nil
}

func (r *AdvertisementPostgres) GetAdvertisementById(id int) (advertisement.AdvertisementDTO, error) {
	var advertisement advertisement.AdvertisementDTO

	query := fmt.Sprint(`
								SELECT a.description,
									   a.title,
									   i.img
									FROM advertisement a
									join advertisement_img ai on a.id = ai.advertisement_id
									join img i on i.id = ai.img_id
									where a.id = $1
									ORDER  BY i.id, a.id
									limit (select count(*) from advertisement)
									`)

	err := r.db.Get(&advertisement, query, id)

	return advertisement, err

}

func (r *AdvertisementPostgres) UpdateAdvertisement(id int, dto advertisement.UpdateAdvertisement) error {
	queryAdv := fmt.Sprint(`
						update advertisement
						set description = $1,
							title       = $2
						where id = $3
						`)

	queryImg := fmt.Sprint(`
						update img
						set img = $1
						where img.id = (SELECT i.id
										FROM advertisement a
												 join advertisement_img ai on a.id = ai.advertisement_id
												 join img i on i.id = ai.img_id
										where a.id = $2
										ORDER BY i.id, a.id
										limit 1)
						`)

	_, err := r.db.Exec(queryAdv, dto.Description, dto.Title, id)

	_, err = r.db.Exec(queryImg, dto.Img, id)

	return err

}
