package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"test-buyer-experience/internal/app/test/getter/testgetter"
	"test-buyer-experience/internal/app/test/sender/testsender"
	"test-buyer-experience/internal/app/test/store/model"
	"test-buyer-experience/internal/app/test/store/teststore"
)

func Test_EmptyDB(t *testing.T) {
	//Arrange
	st := teststore.New()
	getter := testgetter.NewTestGetter()
	sender := testsender.NewTestSender()
	logger, _ := zap.NewDevelopment()
	s := NewServer(st, logger, getter, sender)

	//Act
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodGet,
		"/api/v1/subscription?"+
			"link=https://m.avito.ru/penza/predlozheniya_uslug/uslugi_elektrika._vse_vidy_rabot_1796651105&"+
			"email=elkozlova94@yandex.ru",
		nil,
	)
	s.ServeHTTP(rec, req)

	//Assert
	assert.Equal(t, 200, rec.Code)

	number := "1796651105"
	ad, _ := st.Ads().FindByNumber(number)

	require.EqualValues(t, 1, ad.ID)
	require.Equal(t, number, ad.Number)
	require.Equal(t, "500", ad.Price)

	subscribers, _ := st.Subscribers().FindByAdID(ad.ID)

	require.Equal(t, 1, len(subscribers))
	require.EqualValues(t, 1, subscribers[0].ID)
	require.EqualValues(t, 1, subscribers[0].AdsID)
	require.Equal(t, "elkozlova94@yandex.ru", subscribers[0].Email)
}

func Test_AdditionalSubscriberToSameAd(t *testing.T) {
	//Arrange
	st := teststore.New()
	getter := testgetter.NewTestGetter()
	sender := testsender.NewTestSender()
	logger, _ := zap.NewDevelopment()
	ad := &model.Ads{Number: "1265922529", Price: "220"}
	_ = st.Ads().Save(ad)
	_ = st.Subscribers().Save(&model.Subscribers{AdsID: ad.ID, Email: "test@gmail.com"})
	s := NewServer(st, logger, getter, sender)

	//Act
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodGet,
		"/api/v1/subscription?"+
			"link=https://www.avito.ru/penza/predlozheniya_uslug/uslugi_elektrika_bez_posrednikov_1265922529&"+
			"email=test@yandex.ru",
		nil,
	)
	s.ServeHTTP(rec, req)

	//Assert
	assert.Equal(t, 200, rec.Code)

	number := "1265922529"
	ads, _ := st.Ads().FindAll()

	require.EqualValues(t, len(ads), 1)
	require.EqualValues(t, ads[0].ID, 1)
	require.Equal(t, ads[0].Number, number)
	require.Equal(t, ads[0].Price, "220")

	subscribers, _ := st.Subscribers().FindByAdID(ad.ID)
	require.Equal(t, 2, len(subscribers))

	subscriber1, _ := st.Subscribers().FindByEmailAndAdID(ad.ID, "test@gmail.com")
	require.EqualValues(t, 1, subscriber1.ID)
	require.EqualValues(t, 1, subscriber1.AdsID)
	require.Equal(t, "test@gmail.com", subscriber1.Email)

	subscriber2, _ := st.Subscribers().FindByEmailAndAdID(ad.ID, "test@yandex.ru")
	require.EqualValues(t, 2, subscriber2.ID)
	require.EqualValues(t, 1, subscriber2.AdsID)
	require.Equal(t, "test@yandex.ru", subscriber2.Email)
}

func Test_SendAlert(t *testing.T) {
	//Arrange
	st := teststore.New()
	getter := testgetter.NewTestGetter()
	sender := testsender.NewTestSender()
	logger, _ := zap.NewDevelopment()
	s := NewServer(st, logger, getter, sender)

	number := "909388249"
	oldPrice := "150"
	email := "test@gmail.com"
	newPrice := "500"
	ad := &model.Ads{Number: number, Price: oldPrice}
	_ = st.Ads().Save(ad)
	_ = st.Subscribers().Save(&model.Subscribers{AdsID: ad.ID, Email: email})

	//Act
	s.sendAlert()

	//Assert
	require.Equal(t, number, sender.Number)
	require.Equal(t, oldPrice, sender.OldPrice)
	require.Equal(t, newPrice, sender.NewPrice)
	require.Equal(t, email, sender.Email)

	ad, _ = st.Ads().FindByNumber(number)
	require.Equal(t, newPrice, ad.Price)
}
