package cars_test

import (
	"testing"

	"github.com/matryer/is"

	. "go-bits/variable-names/cars"
)

func Test_Cars_FilterSportscars(t *testing.T) {
	// arrange
	shouldBe := is.New(t)
	cars := NewCars()
	cars.Add(NewCar("1", "Porsche", 250, "Red"))
	cars.Add(NewCar("2", "Ferrari", 250, "Red"))
	cars.Add(NewCar("3", "Porsche", 190, "Red"))
	cars.Add(NewCar("4", "Porsche", 210, "Black"))

	// act
	sportscars := cars.FilterSportscars()

	// assert
	shouldBe.True(len(sportscars) == 1)
}

func Test_Cars_PaintItRed(t *testing.T) {
	// arrange
	shouldBe := is.New(t)
	cars := NewCars()
	cars.Add(NewCar("1", "Porsche", 250, "Red"))
	cars.Add(NewCar("2", "Ferrari", 250, "Red"))
	cars.Add(NewCar("3", "Porsche", 190, "Red"))
	cars.Add(NewCar("4", "Porsche", 210, "Black"))

	// act
	car, err := cars.WithID("4")
	shouldBe.NoErr(err)
	car.PaintItRed()

	// assert
	sportscars := cars.FilterSportscars()
	shouldBe.True(len(sportscars) == 2)
}
