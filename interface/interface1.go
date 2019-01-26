// Go in action
// @jeffotoni
// 2019-01-24

package main

import "fmt"

type Blues struct {
  Volume     int
  Mass       int
  Wavespeed  float64 // wave speed
  Wavelength float64 // length of the wave.
}

type I interface {
  Density() int
  Frequency() float64
}

func (r *Blues) Density() int {
  return r.Mass / r.Volume
}

// f = V / λ
// λ represents the wavelength.
func (f *Blues) Frequency() float64 {
  return f.Wavespeed / f.Wavelength
}

type BossaNova struct {
  Volume     int
  Mass       int
  Wavespeed  float64 // wave speed
  Wavelength float64 // length of the wave.
}

func (r *BossaNova) Density() int {
  return r.Mass / r.Volume
}

// f = V / λ
// λ represents the wavelength.
func (f *BossaNova) Frequency() float64 {
  return f.Wavespeed / f.Wavelength
}

func Calculate(i I) string {

  return fmt.Sprintf("Density: %d\nFrequency: %f", i.Density(), i.Frequency())
}

func main() {

  // A way to use Interface
  var Interf I
  b := Blues{2, 22, 323, 156}
  Interf = &b

  fmt.Println("################################")
  fmt.Println("Blues")
  fmt.Println("Density: ", Interf.Density())
  fmt.Println("Frequency: ", Interf.Frequency())

  bossa := BossaNova{5, 25, 99, 88}
  Interf = &bossa
  fmt.Println("################################")
  fmt.Println("Bossa Nova")
  fmt.Println("Density: ", Interf.Density())
  fmt.Println("Frequency: ", Interf.Frequency())

  // Second way to access interface
  fmt.Println("################################")
  fmt.Println("Blues")
  fmt.Println(Calculate(&b))
  fmt.Println("################################")
  fmt.Println("Bossa")
  fmt.Println(Calculate(&bossa))
  fmt.Println("################################")
}
