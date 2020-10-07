package teststore

import (
	"test-buyer-experience/internal/app/test/store/model"
)

type AdsRepo struct {
	ads map[int32]*model.Ads

	counter int32
}

func (r *AdsRepo) FindAll() ([]*model.Ads, error) {
	ads := make([]*model.Ads, 0)
	for _, v := range r.ads {
		ads = append(ads, v)
	}
	return ads, nil
}

func (r *AdsRepo) FindByNumber(number string) (*model.Ads, error) {
	for _, v := range r.ads {
		if v.Number == number {
			return v, nil
		}
	}
	return nil, nil
}

func (r *AdsRepo) Update(ad *model.Ads) error {
	r.counter++
	r.ads[r.counter] = ad
	return nil
}

func (r *AdsRepo) Save(ad *model.Ads) error {
	r.counter++
	ad.ID = r.counter
	r.ads[r.counter] = ad
	return nil
}
