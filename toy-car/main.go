package main

import "fmt"

type Car struct {
	speed   int
	battery int
}

func NewCar(speed, battery int) *Car {
	car := Car{speed: speed, battery: battery}
	return &car
}
func GetSpeed(car *Car) int {
	return car.speed
}
func GetBattery(car *Car) int {
	return car.battery
}
func ChargeCar(car *Car, minutes int) {
	for i := 0; i < minutes/2; i++ {
		if car.battery < 100 {
			car.battery += 1
		}
	}
}
func TryFinish(car *Car, distance int) string {
	time := float64(distance) / float64(car.speed)
	batteryDrain := float64(distance) / 2
	if car.battery < int(batteryDrain) {
		car.battery = 0
		return ""
	}
	car.battery -= int(batteryDrain)

	return fmt.Sprintf("%.2f", time)
}
