package p_01603_parking_system

type ParkingSystem struct {
	big int
	medium int
	small int
}


func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}


func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 {
		if this.big == 0 {
			return false
		}
		this.big--
		return true
	} else if carType == 2 {
		if this.medium == 0 {
			return false
		}
		this.medium--
		return true
	} else if carType == 3 {
		if this.small == 0 {
			return false
		}
		this.small--
		return true
	}
	return false
}
