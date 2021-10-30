package main

import (
	"fmt"
	"math"
	"strconv"
)

type Movie struct {
	Name string
	Rating float64
}

func (m *Movie) summary() string  {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + ", " + r
}

type Sphere struct {
	Radius float64
}

func(s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}

func(s *Sphere) Volume() float64 {
	radiusCubed := s.Radius * s.Radius * s.Radius
	return (float64(4) / float64(3)) * math.Pi * radiusCubed
}




func main() {
	fmt.Println("METHODS!")

	m := Movie{
		Name: "Rambo",
		Rating: 5.138,
	}

	fmt.Println(m.summary())

	s := Sphere{
		Radius: 5,
	}

	fmt.Println(s.SurfaceArea())
	fmt.Println(s.Volume())

	fmt.Println("***** INTERFACES!!! *****")
}
