package teststore

import (
	"test-buyer-experience/internal/app/test/store"
	"test-buyer-experience/internal/app/test/store/model"
)

type Store struct {
	adsRepo         *AdsRepo
	subscribersRepo *SubscribersRepo
}

func New() *Store {
	return &Store{
		adsRepo: &AdsRepo{
			ads: make(map[int32]*model.Ads, 0),
		},
		subscribersRepo: &SubscribersRepo{
			subscribers: make(map[int32]*model.Subscribers, 0),
		},
	}
}

func (s *Store) Ads() store.AdsRepo {
	return s.adsRepo
}

func (s *Store) Subscribers() store.SubscribersRepo {
	return s.subscribersRepo
}
