package main

import (
	"fmt"
)

func MinimoNotas(value int) (int, map[int]int) {
	notasDisponiveis := []int{200, 100, 50, 20, 10, 5, 2}
	numNotas := make(map[int]int)

	for _, nota := range notasDisponiveis {
		if value >= nota {
			count := value / nota
			numNotas[nota] = count
			value -= count * nota
		}
	}

	return value, numNotas
}

func main() {
	var inputValue int
	var ctrl = 1
	for ctrl == 1 {
		fmt.Print("Digite o valor desejado: ")
		_, err := fmt.Scan(&inputValue)
		if err != nil {
			fmt.Println("Entrada inválida")
		}

		remainingValue, notas := MinimoNotas(inputValue)
		if remainingValue > 0 {
			fmt.Println("O valor não esta disponivel para saque")
		} else {
			fmt.Println("Notas escolhidas:")
			for denomination, count := range notas {
				if count > 0 {
					fmt.Printf("%d notas de %d\n", count, denomination)
				}
			}
		}

		fmt.Println("Deseja fazer outra operação? \n Digite: \n 1 - Sim \n 0 - Não")
		fmt.Scan(&ctrl)
	}

}
