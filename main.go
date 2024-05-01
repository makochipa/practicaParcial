package main

import "github.com/untref-ayp2/data-structures/stack"

func NumeroMaximo(numeros []int) int {
	maximo := 0
	pila := stack.New[int]()
	for _, num := range numeros {
		if pila.IsEmpty() || num > pila.Top() {
			pila.Push(num)
		}
		if num > maximo {
			maximo = num
		}
	}
	return maximo
}