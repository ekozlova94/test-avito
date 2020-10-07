package prodgetter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
	"test-buyer-experience/internal/pkg/forms"
)

type GetterImpl struct {
	cfg *viper.Viper
}

func NewGetter(cfg *viper.Viper) *GetterImpl {
	return &GetterImpl{
		cfg: cfg,
	}
}

func (s *GetterImpl) GetPrice(number string) (string, error) {
	client := &http.Client{}
	response, err := client.Get(
		"https://m.avito.ru/api/14/items/" + number + "?key=" + s.cfg.GetString("key"),
	)
	if err != nil {
		return "", err
	}
	//noinspection GoUnhandledErrorResult
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received not ok response code: %d", response.StatusCode)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	price := forms.Price{}
	if err = json.Unmarshal(data, &price); err != nil {
		return "", err
	}
	if price.Body == nil {
		return "", fmt.Errorf("error while requesting price: %v", price.Body)
	}
	if price.Body.Value == "" {
		return "", fmt.Errorf("unable to get ad price: %v", price.Body.Value)
	}
	return price.Body.Value, nil
}
