package main

import (
	"errors"
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

type Triangle struct {
	base float64
	height float64
}

func(t *Triangle) area() float64 {
	return 0.5 * (t.base * t.height)
}

func(t *Triangle) changeBase(f float64) {
	t.base = f
	return
}

type Robot interface {
	PowerOn() error
}

func Boot(r Robot) error {
	return r.PowerOn()
}

type T850 struct {
	Name string
}

func (a *T850) PowerOn() error  {
	return nil
}

type R2D2 struct {
	Broken bool
}

func(r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
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

	t := Triangle{
		base: 3,
		height: 1,
	}

	fmt.Println(t.area())

	t.changeBase(4)

	fmt.Println(t.area())

	fmt.Println("***** INTERFACES!!! *****")

	terminator := T850{
		Name: "Terminator",
	}

	r2d2 := R2D2{
		Broken: true,
	}

	err := Boot(&terminator)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}

	err = Boot(&r2d2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}


}
