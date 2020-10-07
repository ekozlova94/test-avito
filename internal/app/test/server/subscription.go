package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"test-buyer-experience/internal/app/test/getter"
	"test-buyer-experience/internal/app/test/sender"
	"test-buyer-experience/internal/app/test/store"
	"test-buyer-experience/internal/app/test/store/model"
	"test-buyer-experience/internal/app/test/utils"
)

type server struct {
	Router *gin.Engine
	Store  store.Store
	Log    *zap.Logger
	Getter getter.Getter
	Sender sender.Sender
}

func NewServer(store store.Store, logger *zap.Logger, getter getter.Getter, sender sender.Sender) *server {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	s := &server{
		Router: r,
		Store:  store,
		Log:    logger,
		Getter: getter,
		Sender: sender,
	}
	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.Router.GET("/api/v1/subscription", s.GetLinkAndEmail)
}

func (s *server) GetLinkAndEmail(c *gin.Context) {
	link := c.Query("link")
	email := c.Query("email")

	if !utils.IsEmailValid(email) {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("%s is incorrect", email))
		return
	}

	number, err := utils.IsURLValid(link)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("%s", err))
		return
	}

	if err := s.SetLinkAndEmail(number, email); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}
}

func (s *server) SetLinkAndEmail(number, email string) error {
	ad, err := s.Store.Ads().FindByNumber(number)
	if err != nil {
		return err
	}
	if ad == nil {
		price, err := s.Getter.GetPrice(number)
		if err != nil {
			return err
		}
		ad = &model.Ads{
			Number: number,
			Price:  price,
		}
		if errAd := s.Store.Ads().Save(ad); errAd != nil {
			return errAd
		}
	}
	subscriber, errSub := s.Store.Subscribers().FindByEmailAndAdID(ad.ID, email)
	if errSub != nil {
		return errSub
	}
	if subscriber != nil {
		return nil
	}
	subscriber = &model.Subscribers{AdsID: ad.ID, Email: email}
	if err := s.Store.Subscribers().Save(subscriber); err != nil {
		return err
	}
	return nil
}

func (s *server) BackgroundTask() {
	go s.startCheckChangePrice()
}

func (s *server) startCheckChangePrice() {
	for {
		s.sendAlert()
		time.Sleep(24 * time.Hour)
	}
}

func (s *server) sendAlert() {
	ads, err := s.Store.Ads().FindAll()
	if err != nil {
		s.Log.Error("Error while requesting all ads", zap.Error(err))
	}
	for _, ad := range ads {
		oldPrice := ad.Price
		newPrice, err := s.Getter.GetPrice(ad.Number)
		if err != nil {
			s.Log.Error("Error while requesting new price", zap.Error(err))
			continue
		}
		if oldPrice != newPrice {
			subscribers, err := s.Store.Subscribers().FindByAdID(ad.ID)
			if err != nil {
				s.Log.Error("Error while requesting subscribers", zap.Error(err))
				continue
			}
			for _, subscriber := range subscribers {
				s.Sender.Send(subscriber.Email, ad.Number, oldPrice, newPrice)
			}
			ad.Price = newPrice
			if err := s.Store.Ads().Update(ad); err != nil {
				s.Log.Error("Error while saving ad", zap.Error(err))
				continue
			}
		}
	}
}
