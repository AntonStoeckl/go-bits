package cars

import (
	"errors"
)

// For the sake of this example I simply use alias types for the of Car

type ID string
type CarMake string
type Horsepower uint8
type Color string

const sportscarsMinHorsepower = Horsepower(200)
const sportscarsMake = CarMake("Porsche")
const sportscarsColor = Color("Red")

// Car

type Car struct {
	ID         ID
	CarMake    CarMake
	Horsepower Horsepower
	Color      Color
}

func NewCar(id ID, carMake CarMake, horsepower Horsepower, color Color) *Car {
	return &Car{ID: id, CarMake: carMake, Horsepower: horsepower, Color: color}
}

func (c *Car) PaintItRed() {
	c.Color = "Red"
}

// Cars

type Cars struct {
	items map[ID]*Car
}

func NewCars() *Cars {
	return &Cars{
		make(map[ID]*Car),
	}
}

func (c *Cars) Add(car *Car) {
	if _, alreadyExists := c.items[car.ID]; !alreadyExists {
		c.items[car.ID] = car
	}
}

func (c *Cars) Save(car *Car) error {
	if _, exists := c.items[car.ID]; !exists {
		return errors.New("car does not exist")
	}

	c.items[car.ID] = car

	return nil
}

func (c *Cars) WithID(id ID) (*Car, error) {
	car, found := c.items[id]
	if !found {
		return nil, errors.New("car not found")
	}

	return car, nil
}

func (c *Cars) FilterSportscars() []*Car {
	var sportscars []*Car

	for id, car := range c.items {
		if car.Horsepower > sportscarsMinHorsepower && car.CarMake == sportscarsMake && car.Color == sportscarsColor {
			sportscars = append(sportscars, car)
		}

		// let's just imagine we do something useful with the slice index
		_ = id
	}

	return sportscars
}

func (c *Cars) FilterSportscarsV2() []*Car {
	var sportscars []*Car

	for id, car := range c.items {
		hasEnoughHorsepower := car.Horsepower > sportscarsMinHorsepower
		hasSportyMake := car.CarMake == sportscarsMake
		hasASportyColor := car.Color == sportscarsColor

		isASportscar := hasEnoughHorsepower && hasSportyMake && hasASportyColor

		if isASportscar {
			sportscars = append(sportscars, car)
		}

		// let's just imagine we do something useful with the map index
		_ = id
	}

	return sportscars
}
