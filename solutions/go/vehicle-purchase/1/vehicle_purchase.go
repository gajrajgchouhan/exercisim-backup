package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if (kind == "car" || kind == "truck") {
        return true
    }
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
    var r string;
	if (option1> option2) {
        r = option2
    } else {
        r= option1
    }
	return r + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    var p float64
	if (age < 3) {
        p = 0.8
    } else if (age >= 10) {
    	p = 0.5
    } else {
    	p = 0.7
    }
	return float64(originalPrice) * p
}
