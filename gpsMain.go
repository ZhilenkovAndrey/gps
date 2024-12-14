package main

import (
	. "fmt"
	. "math"
)

type gps struct {
	current     location
	destination location
	world
}

type location struct {
	name      string
	lat, long float64
}

type world struct {
	radius float64
}

type rover struct {
	name string
	gps
}

func (l location) description() string {
	return Sprintf("Station %s - longitud: %f, latitude: %f", l.name, l.long, l.lat)
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := Sincos(rad(p1.lat))
	s2, c2 := Sincos(rad(p2.lat))
	clong := Cos(rad(p1.long - p2.long))
	return w.radius * Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * Pi / 180
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string {
	return Sprintf("%1fkm live to %v", g.distance(), g.destination.description())
}

func main() {
	mars := world(3389.5)
	bradbury := location("Bradbury Landing", -4.5895, 137.4417)
	elysyum := location("elysium Planitia", 4.5, 135.9)

	gps := gps{mars, bradbury, elysyum}

	curiosity := rover{"ROVER", gps}

	Println(curiosity.message())
}
