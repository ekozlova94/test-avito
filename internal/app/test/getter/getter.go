package getter

type Getter interface {
	GetPrice(number string) (string, error)
}
