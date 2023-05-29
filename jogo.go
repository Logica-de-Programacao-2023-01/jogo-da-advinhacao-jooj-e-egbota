package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	jogarNovamente := true

	var tentativasTotal [][]int

	for jogarNovamente {
		fmt.Println(" Jogo da Adivinhação, tente acertar um número de 1 a 100")

		numeroAleatorio := rand.Intn(100) + 1
		tentativas := []int{}
		acertou := false

		for !acertou {
			var x int
			fmt.Print("Digite um número : ")
			fmt.Scanln(&x)

			tentativas = append(tentativas, x)

			if x == numeroAleatorio {
				fmt.Println("Parabéns, você acertou!")
				acertou = true

			} else if x < numeroAleatorio {

				fmt.Println("Tente mais alto dessa vez rsrs")

			} else {
				fmt.Println("Tente mais baixo dessa vez rsrs ")

			}
		}

		tentativasTotal = append(tentativasTotal, tentativas)

		fmt.Print("Deseja jogar novamente? (s/n): ")
		var resposta string
		fmt.Scanln(&resposta)

		if resposta != "s" && resposta != "S" {
			jogarNovamente = false
		}

		fmt.Println()
	}

	fmt.Println("Estatísticas do Jogo ")
	for i, tentativas := range tentativasTotal {
		fmt.Printf("Jogada %d - Tentativas: %d\n", i+1, len(tentativas))
	}
}
