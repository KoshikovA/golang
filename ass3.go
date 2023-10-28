package main

import (
	"fmt"
	"sync"
)

//Command pattern

type CustomerService struct {
    name string
}

type Button struct {
	label       string
	command     Command
}

type AddCustomerCommand struct {
	customerService CustomerService 
}

type Command interface {
	execute()
}



func (b *Button) Click() {
	b.command.execute()
}

func (b *Button) GetLabel() string {
	return b.label
}

func (b *Button) SetLabel(label string) {
	b.label = label
}


func (cs *CustomerService) AddCustomer() {
    fmt.Println("Customer added!")
}

func (ac *AddCustomerCommand) execute() {
	ac.customerService.AddCustomer()
}


//Singleton pattern

type MainCharacter struct {
	Name   string
	Level  int
	Health int
}

type Game struct {
	mainCharacter MainCharacter
	mu            sync.RWMutex
}

var instance *Game
var once sync.Once

func GetGame() *Game {
	once.Do(func() {
		instance = &Game{}
	})
	return instance
}

func (g *Game) CreateMainCharacter(name string, level, health int) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.mainCharacter = MainCharacter{Name: name, Level: level, Health: health}
}

func (g *Game) GetMainCharacter() MainCharacter {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return g.mainCharacter
}

func main() {
	//Command
	service := &CustomerService{"Kaspersky"}
	command := &AddCustomerCommand{customerService: *service}
	button := &Button{"Click here",command}
	button.Click()

    //Singleton
	game := GetGame()
	game.CreateMainCharacter("Hero", 5, 100)
	mainCharacter := game.GetMainCharacter()
	fmt.Printf("Main Character: %s (Level %d, Health %d)\n", mainCharacter.Name, mainCharacter.Level, mainCharacter.Health)
}
