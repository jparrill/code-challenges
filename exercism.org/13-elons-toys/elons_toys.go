package elon

import "fmt"

/*
type Car struct {
	speed        int
	batteryDrain int

	battery  int
	distance int
}
*/


// TODO: define the 'Drive()' method
func (c *Car) Drive() {
    if c.batteryDrain <= c.battery {
    	c.battery -= c.batteryDrain
    	c.distance += c.speed
    }
}

// TODO: define the 'DisplayDistance() string' method
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
    howFar := (c.battery / c.batteryDrain) * c.speed

    if trackDistance <= howFar {
        return true
    }

    return false
}
