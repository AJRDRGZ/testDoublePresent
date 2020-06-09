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