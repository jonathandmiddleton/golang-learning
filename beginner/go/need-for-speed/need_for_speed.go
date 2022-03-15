package speed

type Car struct {
	speed        int
	batteryDrain int
	distance     int
	battery      int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	newCar := NewCar(speed, batteryDrain)
	return newCar
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	newTrack := NewTrack(distance)
	return newTrack
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	newDistance := car.distance + car.speed
	canDrive := (newDistance / car.batteryDrain) - car.battery
	return Car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	car := car.battery / car.batteryDrain

}
