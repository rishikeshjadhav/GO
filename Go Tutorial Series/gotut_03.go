package main

import "fmt"

type car struct {
	gas_pedal      uint16
	brake_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

func main() {
	fmt.Println("Hello from Go programming")

	a_car := car{gas_pedal: 22341, brake_pedal: 0, steering_wheel: 2425, top_speed_kmh: 250.4}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.brake_pedal)
	fmt.Println(a_car.steering_wheel)
	fmt.Println(a_car.top_speed_kmh)
}
