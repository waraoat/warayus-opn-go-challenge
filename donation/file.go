package donation

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"tamboon/go-challenge/cipher"
)

func readCSVFile(filename string) ([]Donation, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Print("Error opening file")
		return nil, err
	}
	defer file.Close()

	rot128Reader, err := cipher.NewRot128Reader(file)
	if err != nil {
		fmt.Print("Error creating rot128 reader")
		return nil, err
	}

	reader := csv.NewReader(rot128Reader)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Print("Error reading file")
		return nil, err
	}

	donations := []Donation{}
	for index, record := range records {
		if index == 0 {
			continue
		}
		
		amountSubunits, _ := strconv.Atoi(record[1])
		expMonth, _ := strconv.Atoi(record[4])
		expYear, _ := strconv.Atoi(record[5])

		donation := Donation{
			Name:           record[0],
			AmountSubunits: amountSubunits,
			CCNumber:       record[2],
			CVV:            record[3],
			ExpMonth:       expMonth,
			ExpYear:        expYear,
		}

		donations = append(donations, donation)
	}

	return donations, nil
}