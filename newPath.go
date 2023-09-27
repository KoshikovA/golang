package main

import (
	"fmt"
)

//Observer
type Worker struct {
	name       string
	hourlyWage float32
	monthHours []int
}

type Employer struct {
	name    string
	workers []*Worker
}

type IWorker interface {
	Work(hours int)
	Update()
}

type IEmployer interface {
	Hire(w *Worker)
	PaySalary()
	Fire(w *Worker)
	NotifyWorkers()
}

func (w *Worker) Work(hours int) {
	fmt.Printf("%s worked %d hours today!\n", w.name, hours)
	w.monthHours = append(w.monthHours, hours)
}

func (w *Worker) Update() {
	fmt.Printf("Hourly Worker %s: Hourly wage $%.2f\n", w.name, w.hourlyWage)
	sum := 0
	for _, h := range w.monthHours {
		sum += h
	}
	fmt.Printf("Hourly Worker %s: Earned $%.2f today \n", w.name, float32(sum)*w.hourlyWage)
}

func (e *Employer) Hire(w *Worker) {
	e.workers = append(e.workers, w)
	fmt.Printf("%s hired to %s company!\n", w.name, e.name)
}

func (e *Employer) PaySalary() {
	e.NotifyWorkers()
	fmt.Println("Company: Paying salaries to all workers.")
}

func (e *Employer) Fire(w *Worker) {
	for i, value := range e.workers {
		if value == w {
			e.workers = append(e.workers[:i], e.workers[i+1:]...)
			fmt.Printf("%s fired from %s company!\n", w.name, e.name)
			break
		}
	}
}

func (e *Employer) NotifyWorkers() {
	for _, worker := range e.workers {
		worker.Update()
	}
}



//Strategy Pattern
type ShippingStrategy interface {
    CalculateCost(distance float64, weight float64) float64
}
type StandardShippingStrategy struct{}

func (s *StandardShippingStrategy) CalculateCost(distance float64, weight float64) float64 {
    ratePerMile := 0.15
    return ratePerMile * distance * weight
}

type ExpressShippingStrategy struct{}

func (e *ExpressShippingStrategy) CalculateCost(distance float64, weight float64) float64 {
    ratePerMile := 0.3
    return ratePerMile * distance * weight
}

type ShippingCostCalc struct {
    strategy ShippingStrategy
}

func (c *ShippingCostCalc) SetStrategy(strategy ShippingStrategy) {
    c.strategy = strategy
}

func (c *ShippingCostCalc) Calculate(distance float64, weight float64) float64 {
    return c.strategy.CalculateCost(distance, weight)
}


func main() {
	a := &Worker{name: "Islam", hourlyWage: 8.73}
	b := &Worker{name: "Dosymzhan", hourlyWage: 6.70}
	c := &Worker{name: "Bob", hourlyWage: 25.00}
	employer := &Employer{name: "Google"}

	employer.Hire(a)
	employer.Hire(b)
	employer.Hire(c)

	a.Work(7)
	b.Work(9)
	c.Work(8)

	employer.Fire(b)

	employer.PaySalary()






	distance := 630.0 
    weight := 5.5    

    calculator := ShippingCostCalc{}

    
    calculator.SetStrategy(&StandardShippingStrategy{})
    standardCost := calculator.Calculate(distance, weight)
    fmt.Printf("\nStandard Shipping Cost: $%.2f\n", standardCost)

    calculator.SetStrategy(&ExpressShippingStrategy{})
    expressCost := calculator.Calculate(distance, weight)
    fmt.Printf("Express Shipping Cost: $%.2f\n", expressCost)
}