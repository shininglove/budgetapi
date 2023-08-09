package my_csv 

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

// "Status Date Original Description Split Type Category Currency Amount User Description Memo Classification Account Name Simple Description"
type purchases struct {
	status              string
	date                time.Time
	originalDescription string
	splitType           string
	category            string
	currency            string
	amount              string
	userDescription     string
	memo                string
	classification      string
	accountName         string
	simpleDescription   string
}

func main() {
	file, err := os.Open("finances_data.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	csvInput, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	data := csvInput[1:]

	dateFormat := "01/02/2006"
	for _, row := range data {
		rowDate, err := time.Parse(dateFormat, row[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(purchases{status: row[0], date: rowDate, originalDescription: row[2], splitType: row[3], category: row[4], currency: row[5], amount: row[6], userDescription: row[7], memo: row[8], classification: row[9], accountName: row[10], simpleDescription: row[11]})
	}

}
