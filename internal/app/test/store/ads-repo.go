package store

import "test-buyer-experience/internal/app/test/store/model"

type AdsRepo interface {
	FindAll() ([]*model.Ads, error)
	FindByNumber(number string) (*model.Ads, error)
	Update(ad *model.Ads) error
	Save(ad *model.Ads) error
}
