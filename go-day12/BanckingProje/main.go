package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"
)

type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Balance     int    `json:"balance"`
	CreditScore int    `json:"creditScore"`
}

type LoanXML struct {
	XMLName  xml.Name `xml:"loan"`
	Customer string   `xml:"customer"`
	Amount   int      `xml:"amount"`
	Status   string   `xml:"status"`
}

func main() {

	data, err := os.ReadFile("customers.json")

	if err != nil {
		panic(err)
	}

	var customers []Customer

	json.Unmarshal(data, &customers)

	customer := customers[0]

	document := fmt.Sprintf(
		"Loan Contract\n\nCustomer: %s\nBalance: %d\nCredit Score: %d\nLoan Amount: 3000 AZN\nStatus: Approved",
		customer.Name,
		customer.Balance,
		customer.CreditScore,
	)

	os.WriteFile("documents/contract_1.txt", []byte(document), 0644)

	logMessage := fmt.Sprintf(
		"%s Loan generated for %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		customer.Name,
	)

	file, _ := os.OpenFile("logs/audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer file.Close()

	file.WriteString(logMessage)

	xmlData := LoanXML{
		Customer: customer.Name,
		Amount:   3000,
		Status:   "Approved",
	}

	output, _ := xml.MarshalIndent(xmlData, "", "  ")

	os.WriteFile("xml/response.xml", output, 0644)

	source, _ := os.Open("documents/contract_1.txt")

	defer source.Close()

	destination, _ := os.Create("backups/backup_contract_1.txt")

	defer destination.Close()

	io.Copy(destination, source)

	fmt.Println("System Completed Successfully")
}
