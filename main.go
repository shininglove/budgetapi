package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"log"
	"os"
	"time"
)

type DateTime struct {
	time.Time
}

func (date *DateTime) MarshalCSV() (string, error) {
	return date.Time.Format("01/02/2006"), nil
}

func (date *DateTime) String() string {
	return date.String()
}

func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("01/02/2006", csv)
	return err
}

// "Status Date Original Description Split Type Category Currency Amount User Description Memo Classification Account Name Simple Description"
type purchases struct {
	Status              string   `csv:"Status"`
	Date                DateTime `csv:"Date"`
	OriginalDescription string   `csv:"Original Description"`
	SplitType           string   `csv:"Split Type"`
	Category            string   `csv:"Category"`
	Currency            string   `csv:"Currency"`
	Amount              string   `csv:"Amount"`
	UserDescription     string   `csv:"User Description"`
	Memo                string   `csv:"Memo"`
	Classification      string   `csv:"Classification"`
	AccountName         string   `csv:"Account Name"`
	SimpleDescription   string   `csv:"Simple Description"`
}

func main() {
	file, err := os.OpenFile("finances_data.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	payments := []*purchases{}

	if err := gocsv.UnmarshalFile(file, &payments); err != nil {
		panic(err)
	}

	for _, payment := range payments {
        fmt.Println("Row: ", payment.Currency, payment.Amount)
	}
	
}
