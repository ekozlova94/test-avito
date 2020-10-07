package testgetter

type GetterTestImpl struct {
	Price string
}

func NewTestGetter() *GetterTestImpl {
	return &GetterTestImpl{
		Price: "500",
	}
}

func (s *GetterTestImpl) GetPrice(_ string) (string, error) {
	return s.Price, nil
}
