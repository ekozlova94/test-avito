package prodsender

import (
	"net/smtp"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type SenderImpl struct {
	cfg *viper.Viper
	log *zap.Logger
}

func NewSender(log *zap.Logger, cfg *viper.Viper) *SenderImpl {
	return &SenderImpl{
		log: log,
		cfg: cfg,
	}
}

func (s *SenderImpl) Send(email, number, oldPrice, newPrice string) {
	from := s.cfg.GetString("email-auth.sender")
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Цена изменилась!\n\n" +
		"Цена, указанная в объявлении №" + number + ", изменена с " + oldPrice + " до " + newPrice

	err := smtp.SendMail(
		s.cfg.GetString("email-auth.host")+":"+s.cfg.GetString("email-auth.port"),
		smtp.PlainAuth(
			"",
			s.cfg.GetString("email-auth.login"),
			s.cfg.GetString("email-auth.password"),
			s.cfg.GetString("email-auth.host"),
		),
		from,
		[]string{to},
		[]byte(msg),
	)
	if err != nil {
		s.log.Error("Error while sending alert", zap.Error(err))
		return
	}
	s.log.Info("Sent alert", zap.String("Subscriber", email))
}
