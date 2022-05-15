package chesgomailing

import (
	"fmt"
	"net/smtp"
)

func Send(config *Config, data DataToSend) error {
	address := fmt.Sprintf("%s:%s", config.Host, config.Port)
	message := []byte(data.Subject + data.Format + data.Body)
	auth := smtp.PlainAuth("", config.From, config.Password, config.Host)
	err := smtp.SendMail(address, auth, config.From, data.To, message)
	if err != nil {
		return err
	}

	return nil
}
