package store

type Store interface {
	Ads() AdsRepo
	Subscribers() SubscribersRepo
}
