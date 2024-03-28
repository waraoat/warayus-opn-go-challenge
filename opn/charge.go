package opn

import (
	"tamboon/logger"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type CreateChargePayload struct {
	Amount      int64
	Card		string
}


func (cfg ClientConfig) CreateCharge (payload CreateChargePayload) (error) {
	client, err := omise.NewClient(cfg.OpnPublicKey, cfg.OpnSecretKey)
	if err != nil {
		logger.ErrorLogger.Printf("Error creating client: %v\n", err)
		return err
	}
	
	charge := &omise.Charge{}

	create := &operations.CreateCharge{
		Amount: payload.Amount,
		Currency: "thb",
		Card: payload.Card,
	}

	if err := client.Do(charge, create); err != nil {
		logger.ErrorLogger.Printf("Error creating charge with payload %+v: %v\n", payload, err)
		return err
	}

	return nil
}