package main

import "fmt"

//Decorator
type Food interface {
	Cost() float64
	Description() string
}

type ConcreteFood struct {
	description string
	cost        float64
}

func (cf ConcreteFood) Cost() float64 {
	return cf.cost
}

func (cf ConcreteFood) Description() string {
	return cf.description
}

type Pizza struct{}

func (p Pizza) Cost() float64 {
	return 10.0
}

func (p Pizza) Description() string {
	return "Pizza"
}

type Pasta struct{}

func (p Pasta) Cost() float64 {
	return 8.0
}

func (p Pasta) Description() string {
	return "Pasta"
}

type Decorator func(Food) Food

func WithExtraCheese(f Food) Food {
	return ConcreteFood{
		description: f.Description() + " with Extra Cheese",
		cost:        f.Cost() + 2.0,
	}
}


func WithExtraToppings(f Food) Food {
	return ConcreteFood{
		description: f.Description() + " with Extra Toppings",
		cost:        f.Cost() + 1.5,
	}
}



//Factory method
type Robot interface {
	Work() string
}

type RobotFactory interface {
	MakeRobot() Robot
}

type WorkerRobot struct {
	name string
}

func (w WorkerRobot) Work() string {
	return fmt.Sprintf("%s is working on a factory.", w.name)
}

type CleanerRobot struct {
	name string
}

func (c CleanerRobot) Work() string {
	return fmt.Sprintf("%s is cleaning the floor.", c.name)
}

type WorkerRobotFactory struct {
	name string
}

func (wf WorkerRobotFactory) MakeRobot() Robot {
	return WorkerRobot{name: wf.name}
}

type CleanerRobotFactory struct {
	name string
}

func (cf CleanerRobotFactory) MakeRobot() Robot {
	return CleanerRobot{name: cf.name}
}


func main() {
	pizza := Pizza{}
	fmt.Printf("Base %s - Cost: $%.2f\n", pizza.Description(), pizza.Cost())

	pizzaWithExtraCheese := WithExtraCheese(pizza)
	fmt.Printf("%s - Cost: $%.2f\n", pizzaWithExtraCheese.Description(), pizzaWithExtraCheese.Cost())

	pasta := Pasta{}
	pastaWithExtraToppings := WithExtraToppings(pasta)
	fmt.Printf("%s - Cost: $%.2f\n\n", pastaWithExtraToppings.Description(), pastaWithExtraToppings.Cost())

	workerRobotFactory := WorkerRobotFactory{name: "Boston Dynamics' Dog"}
	workerRobot := workerRobotFactory.MakeRobot()
	fmt.Println(workerRobot.Work())

	cleanerRobotFactory := CleanerRobotFactory{name: "Samsung vacuum cleaner"}
	cleanerRobot := cleanerRobotFactory.MakeRobot()
	fmt.Println(cleanerRobot.Work())

}
