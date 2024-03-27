package opn

import (
	"fmt"

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
		return err
	}
	
	charge := &omise.Charge{}

	create := &operations.CreateCharge{
		Amount: payload.Amount,
		Currency: "thb",
		Card: payload.Card,
	}

	if err := client.Do(charge, create); err != nil {
		fmt.Printf("Error creating charge: %v\n", err)
		return err
	}

	return nil
}