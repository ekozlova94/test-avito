package sqlstore

import (
	"database/sql"

	"test-buyer-experience/internal/app/test/store/model"
)

type AdsRepo struct {
	db *sql.DB
}

func (r *AdsRepo) FindAll() ([]*model.Ads, error) {
	rows, err := r.db.Query("SELECT * FROM ads")
	if err != nil {
		return nil, err
	}
	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	ads := make([]*model.Ads, 0)
	for rows.Next() {
		ad := new(model.Ads)
		if err := rows.Scan(&ad.ID, &ad.Number, &ad.Price); err != nil {
			return nil, err
		}
		ads = append(ads, ad)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ads, nil
}

func (r *AdsRepo) FindByNumber(number string) (*model.Ads, error) {
	var ad model.Ads
	err := r.db.QueryRow(
		"SELECT * FROM ads WHERE number = $1",
		number,
	).Scan(&ad.ID, &ad.Number, &ad.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &ad, nil
}

func (r *AdsRepo) Update(ad *model.Ads) error {
	_, err := r.db.Exec(
		"UPDATE ads SET price = $2 WHERE id = $1",
		ad.ID,
		ad.Price,
	)
	return err
}

func (r *AdsRepo) Save(ad *model.Ads) error {
	return r.db.QueryRow(
		"INSERT INTO ads (number, price) VALUES ($1, $2) RETURNING id",
		ad.Number,
		ad.Price,
	).Scan(&ad.ID)
}
