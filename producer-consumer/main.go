package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, totalPizzas int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order %d. Please wait...\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		totalPizzas++

		fmt.Printf("Making pizza %d. It will take %d seconds\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza %d\n", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza %d\n", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("*** Pizza order %d is ready!\n", pizzaNumber)
		}

		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	var i = 0

	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	go pizzeria(pizzaJob)

	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order %d is ready!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("Order %d failed!", i.pizzaNumber)
			}
		} else {
			color.Cyan("Done making pizzas!")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("Error closing channel: %v", err)
			}
		}
	}

	color.Cyan("-----------------")
	color.Cyan("Done for the day!")

	color.Cyan("Total pizzas made: %d", totalPizzas)
	color.Cyan("Total pizzas failed: %d", pizzasFailed)
	color.Cyan("Total pizzas made: %d", pizzasMade)
	color.Cyan("-----------------")

	switch {
	case pizzasFailed > 9:
		color.Red("It was a rough day...")
	case pizzasFailed >= 6:
		color.Yellow("It was not a very good day...")
	case pizzasFailed >= 4:
		color.Yellow("It was an OK day...")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day...")
	default:
		color.Green("It was a great day!")
	}
}
