package store

import "test-buyer-experience/internal/app/test/store/model"

type SubscribersRepo interface {
	FindByAdID(id int32) ([]*model.Subscribers, error)
	FindByEmailAndAdID(id int32, email string) (*model.Subscribers, error)
	Save(subscriber *model.Subscribers) error
}
