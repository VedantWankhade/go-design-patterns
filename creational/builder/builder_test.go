package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	buildDirector := BuildDirector{}
	carBuilder := GetVehicleBuilder(VehicleCar)
	buildDirector.SetBuilder(carBuilder)
	car1 := buildDirector.Contruct()
	car2 := buildDirector.Contruct()
	fmt.Println(car1, car2)
	bikeBuilder := GetVehicleBuilder(VehicleBike)
	buildDirector.SetBuilder(bikeBuilder)
	bike1 := buildDirector.Contruct()
	bike2 := buildDirector.Contruct()
	fmt.Println(bike1, bike2)
	fmt.Println(car1 == car2)
	fmt.Println(bike1 == bike2)
}
