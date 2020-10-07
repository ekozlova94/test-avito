package sqlstore

import (
	"database/sql"

	"test-buyer-experience/internal/app/test/store"
)

type Store struct {
	db              *sql.DB
	adsRepo         *AdsRepo
	subscribersRepo *SubscribersRepo
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
		adsRepo: &AdsRepo{
			db: db,
		},
		subscribersRepo: &SubscribersRepo{
			db: db,
		},
	}
}

func (s *Store) Ads() store.AdsRepo {
	return s.adsRepo
}

func (s *Store) Subscribers() store.SubscribersRepo {
	return s.subscribersRepo
}
