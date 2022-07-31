package main

import "fmt"

type Rotina interface {
	IniciarTrabalho()
	PausarTrabalho()
	FinalizarTrabalho()
}

type RotinaHumano interface {
	Rotina
	TomarCafe()
}

type RotinaRobo interface {
	Rotina
	CarregarBateria()
}

type Humano struct{}
type Robo struct{}

func (h *Humano) IniciarTrabalho()   { fmt.Println("iniciar trabalho humano") }
func (h *Humano) PausarTrabalho()    { fmt.Println("pausar trabalho humano") }
func (h *Humano) FinalizarTrabalho() { fmt.Println("finalizar trabalho humano") }
func (h *Humano) TomarCafe()         { fmt.Println("tomar cafe") }

func (r *Robo) IniciarTrabalho()   { fmt.Println("iniciar trabalho robo") }
func (r *Robo) PausarTrabalho()    { fmt.Println("pausar trabalho robo") }
func (r *Robo) FinalizarTrabalho() { fmt.Println("finalizar trabalho robo") }
func (r *Robo) CarregarBateria()   { fmt.Println("carregar bateria robo") }

func main() {
  var ihum RotinaHumano
	var irobo RotinaRobo

	hum := &Humano{}
	robo := &Robo{}

	ihum = hum
	irobo = robo

	ihum.TomarCafe()
	irobo.CarregarBateria()
}
