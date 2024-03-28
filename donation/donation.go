package donation

import (
	"fmt"
	"tamboon/helper"
	"tamboon/opn"
)

type Donation struct {
	Name           string
	AmountSubunits int
	CCNumber       string
	CVV            string
	ExpMonth       int
	ExpYear        int
}

type Summary struct {
	TotalReceived float32
	SuccesfulDonations float32
	FaultyDonations float32
	TotalPerName map[string]float32
}

func Process(filename string) {
	cfg := GetEnv()

	donations, err := readCSVFile(filename)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return
	}

	summary := Summary{
		TotalReceived: 0,
		SuccesfulDonations: 0,
		FaultyDonations: 0,
		TotalPerName: map[string]float32{},
	}

	for index, donation := range donations {
		helper.ShowProgressBar(index, len(donations))
		summary.TotalReceived += float32(donation.AmountSubunits)

		client := opn.ClientConfig{
			OpnPublicKey: cfg.OpnPublicKey,
			OpnSecretKey: cfg.OpnSecretKey,
		}
		
		token, err := client.CreateToken(opn.CreateTokenPayload{
			Name: donation.Name,
			Number: donation.CCNumber,
			ExpirationMonth: donation.ExpMonth,
			ExpirationYear: donation.ExpYear + 1,
			SecurityCode: donation.CVV,
		})
		if err != nil {
			summary.FaultyDonations += float32(donation.AmountSubunits)
			continue
		}

		err = client.CreateCharge(opn.CreateChargePayload{
			Amount: int64(donation.AmountSubunits),
			Card: token.ID,
		})
		if err != nil {
			summary.FaultyDonations += float32(donation.AmountSubunits)
			continue
		}
		
		summary.SuccesfulDonations += float32(donation.AmountSubunits)
		summary.TotalPerName[donation.Name] += float32(donation.AmountSubunits)
	}

	summary.PrintLog()
}

func (s Summary) PrintLog () {
	topThree := []string{}
	for i := 0; i < 3; i++ {
		max := float32(0)
		maxName := ""
		for name, total := range s.TotalPerName {
			if total > max {
				max = total
				maxName = name
			}
		}
		if maxName != "" {
			topThree = append(topThree, maxName)
			delete(s.TotalPerName, maxName)
		}
	}

	var avgPerPerson float32
	if (len(s.TotalPerName) == 0) {
		avgPerPerson = 0
	} else {
		avgPerPerson = s.SuccesfulDonations/float32(len(s.TotalPerName))
	}

	fmt.Printf("done.\n\n")
	fmt.Printf("        total received: THB %.2f\n", s.TotalReceived)
	fmt.Printf("  successfully donated: THB %.2f\n", s.SuccesfulDonations)
	fmt.Printf("       faulty donation: THB %.2f\n", s.FaultyDonations)
	fmt.Printf("    average per person: THB %.2f\n", avgPerPerson)
	fmt.Printf("            top donors:\n")
	for i, donor := range topThree {
		fmt.Printf("                        %d. %s\n", i+1, donor)
	}
}

