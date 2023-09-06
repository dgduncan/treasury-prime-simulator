package main

import (
	"os"
	"sort"
	treasuryprimesandbox "teasury_prime_sandbox"
	"time"

	"github.com/gocarina/gocsv"
)

// type BookRequest struct {
// 	Amount string `json:"amount"`
// 	From   string `json:"from_account_id"`
// 	To     string `json:"to_account_id"`
// }

func main() {

	// c := http.DefaultClient
	// body, err := json.Marshal(BookRequest{
	// 	Amount: "100.00",
	// 	From:   "acct_11jdzn8jnvnt18",
	// 	To:     "acct_11jdzn8jnvnt0n",
	// })

	// if err != nil {
	// 	panic(err)
	// }
	// r, err := http.NewRequest(http.MethodPost, "https://api.sandbox.treasuryprime.com/book", bytes.NewReader(body))
	// if err != nil {
	// 	panic(err)
	// }
	// r.SetBasicAuth("{{TREASURY_PRIME_KEY}}", "{{TREASURY_PRIME_SECRET}}")
	// r.Header.Set("Content-Type", "application/json")

	// resp, err := c.Do(r)
	// if err != nil {
	// 	panic(err)
	// }

	// defer resp.Body.Close()

	// resBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// log.Printf("status %d", resp.StatusCode)
	// log.Printf("res body: %s", string(resBody))

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.File
	file, _ := os.Open("schedule.csv")

	// // Checks for the error
	// if err != nil {
	// 	log.Fatal("Error while reading the file", err)
	// }

	// // Closes the file
	// defer file.Close()

	// // The csv.NewReader() function is called in
	// // which the object os.File passed as its parameter
	// // and this creates a new csv.Reader that reads
	// // from the file
	// reader := csv.NewReader(file)

	// // ReadAll reads all the records from the CSV file
	// // and Returns them as slice of slices of string
	// // and an error if any
	// records, err := reader.ReadAll()

	// // Checks for the error
	// if err != nil {
	// 	fmt.Println("Error reading records")
	// }

	// // Loop to iterate through
	// // and print each of the string slice
	// for _, eachrecord := range records {
	// 	fmt.Println(eachrecord)
	// }

	defer file.Close()

	contrubitons := []*treasuryprimesandbox.ContributionRecord{}

	if err := gocsv.UnmarshalFile(file, &contrubitons); err != nil { // Load clients from file
		panic(err)
	}

	// now := time.Now()
	sort.Slice(contrubitons, func(i, j int) bool {
		return contrubitons[i].ContributionReceiptDate.Time.Before(contrubitons[j].ContributionReceiptDate.Time)
	})

	s := treasuryprimesandbox.Simulator{}
	s.BegnningOfDay = func() {
		// emit balances
		// emit days transactions
	}
	s.EndofDay = func() {
		// emit amount placed into sweep
	}

	startDate, err := time.Parse("2006-01-02", "2023-01-01")
	if err != nil {
		panic(err)
	}

	endDate, err := time.Parse("2006-01-02", "2023-07-01")
	if err != nil {
		panic(err)
	}

	s.Start = startDate
	s.End = endDate
	s.InterstRate = 0.043

	campaign := treasuryprimesandbox.Campaign{
		Candidate:    "Test Candidate",
		Sweep:        0,
		SweepCeiling: 10000,
		Checking:     0,
		HighYield:    0,
		Contribution: contrubitons,
	}

	s.Begin(&campaign)
	// for _, client := range clients {
	// 	fmt.Println(client.ContributionReceiptDate, client.ContributionReceiptAmount)
	// }

	// if _, err := clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
	// 	panic(err)
	// }

}
