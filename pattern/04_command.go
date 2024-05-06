package main

import (
	"fmt"
)

type Command interface {
	execute()
}

// Receiver
type Restaurant struct {
	TotalDishes   int
	CleanedDishes int
}

func NewResteraunt() *Restaurant {
	const totalDishes = 10
	return &Restaurant{
		TotalDishes:   totalDishes,
		CleanedDishes: totalDishes,
	}
}

// Concrete commands
type MakePizzaCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakePizzaCommand) execute() {
	c.restaurant.CleanedDishes -= c.n
	fmt.Println("made", c.n, "pizzas")
}

type MakeSaladCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakeSaladCommand) execute() {
	c.restaurant.CleanedDishes -= c.n
	fmt.Println("made", c.n, "salads")
}

type CleanDishesCommand struct {
	restaurant *Restaurant
}

func (c *CleanDishesCommand) execute() {
	c.restaurant.CleanedDishes = c.restaurant.TotalDishes
	fmt.Println("dishes cleaned")
}

// Commands constructors
func (r *Restaurant) MakePizza(n int) Command {
	return &MakePizzaCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) MakeSalad(n int) Command {
	return &MakeSaladCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) CleanDishes() Command {
	return &CleanDishesCommand{
		restaurant: r,
	}
}

// Invoker
type Cook struct {
	Commands []Command
}

func (c *Cook) executeCommands() {
	for _, c := range c.Commands {
		c.execute()
	}
}

func main() {
	r := NewResteraunt()

	//commands
	tasks := []Command{
		r.MakePizza(2),
		r.MakeSalad(1),
		r.MakePizza(3),
		r.CleanDishes(),
		r.MakePizza(4),
		r.CleanDishes(),
	}

	//invokers
	cooks := []*Cook{
		&Cook{},
		&Cook{},
	}

	for i, task := range tasks {
		cook := cooks[i%len(cooks)]
		cook.Commands = append(cook.Commands, task)
	}

	for i, c := range cooks {
		fmt.Println("cook", i, ":")
		c.executeCommands()
	}
}
