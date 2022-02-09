package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"

}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	vehicle := option2

	// conditional statement checking order
	if option1 < option2 {
		vehicle = option1
	}

	// return string
	return vehicle
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var price float64

	/*
			1,2 - if one
			3,4,5,6,7,8,9,10 - if two
			10+ - if three

		-- also possible:
			switch{
			case age < 3:
				price = originalPrice * 0.8
			case age > 3 && age < 10:
				price = originalPrice * 0.7
			default:
				price = originalPrice * 0.5
			}
	*/

	if age < 3 {
		price = originalPrice * 0.8
	} else if age < 10 {
		price = originalPrice * 0.7
	} else {
		price = originalPrice * 0.5
	}

	return price
}
