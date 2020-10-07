package sender

type Sender interface {
	Send(email, number, oldPrice, newPrice string)
}
