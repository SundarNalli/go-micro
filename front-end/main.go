package main

import "fmt"

type MotorCycle struct {
	WheelCounts int
	TypeOf      string
	Riders      int
}

func (m *MotorCycle) GetWheels() int {
	return m.WheelCounts
}

func (m *MotorCycle) GetTypeOf() string {
	return m.TypeOf
}

type Car struct {
	WheelCounts int
	TypeOf      string
	Passangers  int
}

func (m *Car) GetWheels() int {
	return m.WheelCounts
}

func (m *Car) GetTypeOf() string {
	return m.TypeOf
}

type Vehicle interface {
	GetWheels() int
	GetTypeOf() string
}

func main() {
	myCycle := MotorCycle{
		WheelCounts: 2,
		TypeOf:      "Motor Cycle",
		Riders:      2,
	}

	bmwCar := Car{
		WheelCounts: 4,
		TypeOf:      "Car",
		Passangers:  5,
	}

	PrintVehicle(&myCycle)
	PrintVehicle(&bmwCar)
}

func PrintVehicle(v Vehicle) {
	fmt.Printf("Type of %s, has %d wheels!", v.GetTypeOf(), v.GetWheels())
}
