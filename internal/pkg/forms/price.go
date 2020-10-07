package forms

type Price struct {
	Body *Body `json:"price"`
}

type Body struct {
	Value string `json:"value"`
}
