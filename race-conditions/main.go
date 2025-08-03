package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	var bankBalance int
	var balance sync.Mutex

	fmt.Printf("Initial Account Balance: $%d.00\n", bankBalance)

	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time Job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount

				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On week %d, your earned $%d.00 from %s.\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()

	fmt.Printf("Final Account Balance: $%d.00\n", bankBalance)
}
