package main

import "fmt"

const ebuK = 373.0

func main(){

	tempK := ebuK
	tempC := (tempK - 273.0)

	fmt.Printf("A temperatura de ebulição da água em °K é de %g e em °C é %g", tempK, tempC)

}