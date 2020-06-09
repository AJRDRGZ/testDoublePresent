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
type mockCapitanAmerica struct {
	cantidadMovimientos int
}

func (s *mockCapitanAmerica) Pelear() {
	s.cantidadMovimientos += s.cantidadMovimientos
	fmt.Println("golpe capitan America del doble") // HL
}

func Test_Accion(t *testing.T) {
	capi := &mockCapitanAmerica{} // HL
	escena := NuevaEscena(capi)

	_ = escena.Accion()

	assert.Equal(t, capi.cantidadMovimientos, 3) // HL
}

//END OMIT
