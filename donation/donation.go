package donation

import (
	"fmt"
	"tamboon/helper"
	opn "tamboon/opn"
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
	counts map[string]int
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
		counts: map[string]int{},
	}

	for _, donation := range donations {
		summary.TotalReceived += float32(donation.AmountSubunits)

		client := opn.ClientConfig{
			OpnPublicKey: cfg.OpnPublicKey,
			OpnSecretKey: cfg.OpnSecretKey,
		}
		
		token, err := client.CreateToken(opn.CreateTokenPayload{
			Name: donation.Name,
			Number: donation.CCNumber,
			ExpirationMonth: donation.ExpMonth,
			ExpirationYear: donation.ExpYear,
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
		summary.counts[donation.Name]++
	}

	summary.PrintLog()
}

func (s Summary) PrintLog () {
	topThree := []string{}
	for i := 0; i < 3; i++ {
		maxCount := 0
		var maxVal string
		for val, count := range s.counts {
			if count > maxCount && !helper.Contains(topThree, val) {
				maxCount = count
				maxVal = val
			}
		}
		if maxVal != "" {
			topThree = append(topThree, maxVal)
		}
	}

	var avgPerPerson float32
	if (len(s.counts) == 0) {
		avgPerPerson = 0
	} else {
		avgPerPerson = s.SuccesfulDonations/float32(len(s.counts))
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

