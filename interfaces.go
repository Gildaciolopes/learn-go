// Interfaces são coleções de métodos.

package main

import (
	"fmt"
	"math" 
)

// aqui temos uma interface usada para formas geométricas.
type geometria interface {
	area() float64
	perimetro() float64
}

// área de um quadrado: lado² ou lado*lado
// perímetro = soma dos lados
type quadrado struct {
	lado float64
}
type círculo struct { //área círculo: (Pi)*raio² perímetro do círculo: 2*r*(pi)
	raio float64
}

// usaremos geometria nos quadrados.
func (q quadrado) area() float64 {
	return q.lado * q.lado
}
func (q quadrado) perimetro() float64 {
	return 4 * q.lado
}

// A implementação dos circulos.
func (c círculo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c círculo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// função genérica medir trabalhando em qualquer geometria.
func medir(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {
	q := quadrado{lado: 25}
	c := círculo{raio: 100}

	medir(q)
	medir(c)
}