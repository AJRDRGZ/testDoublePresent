package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type LibretoCapi interface {
	Pelear()
}

type Escena struct {
	CapitanAmerica LibretoCapi
}

func NuevaEscena(c LibretoCapi) Escena {
	return Escena{CapitanAmerica: c}
}

func (e Escena) Accion() int {
	e.CapitanAmerica.Pelear()
	fmt.Println("golpe soldado del invierno")
	e.CapitanAmerica.Pelear()
	e.CapitanAmerica.Pelear()

	return 5 // minutos
}

//START OMIT
type stubCapitanAmerica struct{}

func (s stubCapitanAmerica) Pelear() {
	fmt.Println("golpe capitan America del doble") // HL
}

func Test_Accion(t *testing.T) {
	capi := stubCapitanAmerica{} // HL
	escena := NuevaEscena(capi)

	duracion := escena.Accion()

	assert.Equal(t, duracion, 5) // HL
}

//END OMIT
