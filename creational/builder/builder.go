package builder

import "fmt"

type VehicleType int

const (
	VehicleCar VehicleType = iota
	VehicleBike
)

type Vehicle interface {
	Drive() string
}

type VechicleBuilder interface {
	New()
	BuildWheels()
	BuildSeats()
	BuildStructure()
	Build() Vehicle
}

type Car struct {
	Structure string
	Wheels    int
	Seats     int
}

func (c *Car) Drive() string {
	return fmt.Sprintf("%v is driving\n", *c)
}

type CarBuilder struct {
	c *Car
}

func GetVehicleBuilder(t VehicleType) VechicleBuilder {
	if t == VehicleCar {
		return &CarBuilder{&Car{}}
	} else if t == VehicleBike {
		return &BikeBuilder{&Bike{}}
	}
	return &CarBuilder{&Car{}}
}

func (b *CarBuilder) New() {
	b.c = &Car{}
}

func (b *CarBuilder) BuildWheels() {
	b.c.Wheels = 4
}

func (b *CarBuilder) BuildSeats() {
	b.c.Seats = 4
}

func (b *CarBuilder) BuildStructure() {
	b.c.Structure = "big"
}

func (b *CarBuilder) Build() Vehicle {
	return b.c
}

type Bike struct {
	Structure string
	Wheels    int
	Seats     int
}

func (c *Bike) Drive() string {
	return fmt.Sprintf("%v is driving\n", *c)
}

type BikeBuilder struct {
	c *Bike
}

func (b *BikeBuilder) New() {
	b.c = &Bike{}
}

func (b *BikeBuilder) BuildWheels() {
	b.c.Wheels = 2
}

func (b *BikeBuilder) BuildSeats() {
	b.c.Seats = 2
}

func (b *BikeBuilder) BuildStructure() {
	b.c.Structure = "small"
}

func (b *BikeBuilder) Build() Vehicle {
	return b.c
}

type BuildDirector struct {
	builder VechicleBuilder
}

func (db *BuildDirector) SetBuilder(b VechicleBuilder) {
	db.builder = b
}

func (bd *BuildDirector) Contruct() Vehicle {
	bd.builder.New()
	bd.builder.BuildSeats()
	bd.builder.BuildWheels()
	bd.builder.BuildStructure()
	return bd.builder.Build()
}
