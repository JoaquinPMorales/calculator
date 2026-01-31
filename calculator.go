// Package calculadora proporciona herramientas aritméticas genéricas.
package calculadora

import "golang.org/x/exp/constraints"

// Number es una interfaz que agrupa tipos numéricos enteros y de punto flotante.
type Number interface {
	constraints.Integer | constraints.Float
}

// Add suma dos números de cualquier tipo que implemente la interfaz Number.
//
// Esta función es genérica y funciona tanto con int, int64, float64, etc.
// Para más información sobre la adición:
// https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
