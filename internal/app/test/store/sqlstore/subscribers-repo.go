package sqlstore

import (
	"database/sql"

	"test-buyer-experience/internal/app/test/store/model"
)

type SubscribersRepo struct {
	db *sql.DB
}

func (r *SubscribersRepo) FindByAdID(id int32) ([]*model.Subscribers, error) {
	rows, err := r.db.Query("SELECT * FROM subscribers WHERE ads_id = $1", id)
	if err != nil {
		return nil, err
	}
	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	subscribers := make([]*model.Subscribers, 0)
	for rows.Next() {
		subscriber := new(model.Subscribers)
		if err := rows.Scan(&subscriber.ID, &subscriber.AdsID, &subscriber.Email); err != nil {
			return nil, err
		}
		subscribers = append(subscribers, subscriber)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return subscribers, nil
}

func (r *SubscribersRepo) FindByEmailAndAdID(id int32, email string) (*model.Subscribers, error) {
	var subscriber model.Subscribers
	err := r.db.QueryRow("SELECT * FROM subscribers WHERE ads_id = $1 AND email = $2", id, email).
		Scan(&subscriber.ID, &subscriber.AdsID, &subscriber.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &subscriber, nil
}

func (r *SubscribersRepo) Save(subscriber *model.Subscribers) error {
	return r.db.QueryRow(
		"INSERT INTO subscribers (ads_id, email) VALUES ($1,$2)",
		subscriber.AdsID,
		subscriber.Email,
	).Err()
}
