package treasuryprimesandbox

import (
	"fmt"
	"time"
)

type Simulator struct {
	Start          time.Time
	End            time.Time
	RealTimePerDay time.Duration
	BegnningOfDay  func()
	EndofDay       func()
	InterstRate    float64
	Vig            float64
}

func (s Simulator) Begin(campaign *Campaign) {

	pd := s.Start
	cd := s.Start
	monthbalances := []float64{}
	dsi := 0
	totalinterest := float64(0)
	for {
		if cd.Equal(s.End) {
			break
		}

		dayscontributions := make([]*ContributionRecord, 0)

		fmt.Printf("----------%s------------\n", cd.Format(time.DateOnly))
		for {
			if dsi < len(campaign.Contribution) && campaign.Contribution[dsi].ContributionReceiptDate.Equal(cd) {
				dayscontributions = append(dayscontributions, campaign.Contribution[dsi])
				dsi++
			} else {
				break
			}
		}
		fmt.Printf("Number of contributions - %d\n", len(dayscontributions))

		for _, v := range dayscontributions {
			campaign.Checking += v.ContributionReceiptAmount
		}

		bsweep := campaign.Checking
		sweepamout := float64(0)
		if campaign.Checking > campaign.SweepCeiling {
			sweepamout = campaign.Checking - campaign.SweepCeiling
			campaign.HighYield += sweepamout
			campaign.Checking -= sweepamout
		}
		monthbalances = append(monthbalances, campaign.HighYield)

		fmt.Printf("Cash in checking before sweep - %.2f\n", bsweep)
		fmt.Printf("Total available to sweep - %.2f\n", sweepamout)
		fmt.Printf("Cash in high yield EOD - %.2f\n", campaign.HighYield)

		fmt.Print("--------------------------------\n\n")

		pd = cd
		cd = cd.AddDate(0, 0, 1)

		if pd.Month() != cd.Month() {
			totalDays := len(monthbalances)
			runningAverage := float64(0)

			for _, v := range monthbalances {
				runningAverage += v
			}
			average := runningAverage / float64(totalDays)

			fmt.Printf("------------------------INTEREST--------------------------\n")
			fmt.Printf("Average Monthly Balance : %.2f\n", average)
			fmt.Printf("Interest earned on average monthly balance at %.3f : %.2f\n", s.InterstRate, average*s.InterstRate)
			fmt.Printf("---------------------------------------------------------\n")
			totalinterest += average * s.InterstRate
			campaign.HighYield += average * s.InterstRate

			monthbalances = []float64{}

		}

		time.Sleep(s.RealTimePerDay)

	}

	pi := float64(0)
	for _, v := range campaign.Contribution {
		if v.ContributorName == "SUNRISE BANKS" {
			pi += v.ContributionReceiptAmount
		}
	}

	fmt.Println("\n----------------------------------------------")
	fmt.Printf("Previous Interest %.2f\n", pi)
	fmt.Printf("Total Interest %.2f\n", totalinterest)

}
