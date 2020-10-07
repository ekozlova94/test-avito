package testsender

type SenderTestImpl struct {
	Email    string
	Number   string
	OldPrice string
	NewPrice string
}

func NewTestSender() *SenderTestImpl {
	return &SenderTestImpl{}
}

func (s *SenderTestImpl) Send(email, number, oldPrice, newPrice string) {
	s.Email = email
	s.Number = number
	s.OldPrice = oldPrice
	s.NewPrice = newPrice
}
