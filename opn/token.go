package opn

import (
	"tamboon/logger"
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type CreateTokenPayload struct {
	Name            string
	Number          string
	ExpirationMonth int
	ExpirationYear  int
	SecurityCode    string
}

func (cfg ClientConfig) CreateToken(payload CreateTokenPayload) (*omise.Token,  error) {
	client, err := omise.NewClient(cfg.OpnPublicKey, cfg.OpnSecretKey)
	if err != nil {
		logger.ErrorLogger.Printf("Error creating client: %v\n", err)
		return nil, err
	}

	token := &omise.Token{}

	create := &operations.CreateToken{
		Name:            payload.Name,
		Number:          payload.Number,
		ExpirationMonth: time.Month(payload.ExpirationMonth),
		ExpirationYear:  payload.ExpirationYear,
		SecurityCode:    payload.SecurityCode,
	}

	if err := client.Do(token, create); err != nil {
		logger.ErrorLogger.Printf("Error creating token with payload %+v: %v\n", payload, err)
		return nil, err
	}

	return token, nil
}
