package teststore

import (
	"test-buyer-experience/internal/app/test/store/model"
)

type SubscribersRepo struct {
	subscribers map[int32]*model.Subscribers

	counter int32
}

func (r *SubscribersRepo) FindByAdID(id int32) ([]*model.Subscribers, error) {
	subscribers := make([]*model.Subscribers, 0)
	for _, v := range r.subscribers {
		if v.AdsID == id {
			subscribers = append(subscribers, v)
		}
	}
	return subscribers, nil
}

func (r *SubscribersRepo) FindByEmailAndAdID(id int32, email string) (*model.Subscribers, error) {
	for _, v := range r.subscribers {
		if v.AdsID == id && v.Email == email {
			return v, nil
		}
	}
	return nil, nil
}

func (r *SubscribersRepo) Save(subscriber *model.Subscribers) error {
	r.counter++
	subscriber.ID = r.counter
	r.subscribers[r.counter] = subscriber
	return nil
}
