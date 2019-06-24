package main

import "fmt"

const usixtennbitmax float64 = 65535
const kmh_multiple float64 = 1.6977

type car struct {
	gas_pedal      uint16
	brake_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixtennbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixtennbitmax / kmh_multiple)
}

func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func newer_top_speed(c car, speed float64) car {
	c.top_speed_kmh = speed
	return c
}

func main() {
	a_car := car{
		gas_pedal:      22314,
		brake_pedal:    0,
		steering_wheel: 111,
		top_speed_kmh:  12.0,
	}
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	a_car.new_top_speed(100)
	// a_car = newer_top_speed(a_car, 100)

	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}
