package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Filme struct {
	Nome     string     `json:"nome"`
	Assentos [][]string `json:"assentos"`
}

type Cinema struct {
	Filmes []Filme `json:"filmes"`
}

const arquivo = "data.json"

func carregar() Cinema {
	var cinema Cinema

	data, err := os.ReadFile(arquivo)
	if err != nil {
		fmt.Println("Erro ao carregar dados")
		return cinema
	}

	json.Unmarshal(data, &cinema)
	return cinema
}

func salvar(cinema Cinema) {
	data, _ := json.MarshalIndent(cinema, "", "  ")
	os.WriteFile(arquivo, data, 0644)
}

func listarFilmes(cinema Cinema) {
	for i, f := range cinema.Filmes {
		fmt.Printf("%d - %s\n", i+1, f.Nome)
	}
	fmt.Println()
}

func mostrarAssentos(filme Filme) {
	fmt.Println("\nAssentos (O = livre | X = ocupado)")
	for i, linha := range filme.Assentos {
		fmt.Printf("%d - ", i)
		for _, col := range linha {
			fmt.Printf("%s ", col)
		}
		fmt.Println()
	}
	fmt.Println()
}

func escolherAssento(filme *Filme) bool {
	mostrarAssentos(*filme)

	var linha, coluna int
	fmt.Print("Linha: ")
	fmt.Scanln(&linha)
	fmt.Print("Coluna: ")
	fmt.Scanln(&coluna)

	if filme.Assentos[linha][coluna] == "X" {
		fmt.Println("Assento já ocupado!\n")
		return false
	}

	filme.Assentos[linha][coluna] = "X"
	fmt.Println("Assento reservado!\n")
	return true
}

func pagamento() bool {
	fmt.Println("\n--- PAGAMENTO ---")
	fmt.Println("1 - Cartão")
	fmt.Println("2 - PIX")

	var op int
	fmt.Print("Escolha: ")
	fmt.Scanln(&op)

	if op == 1 || op == 2 {
		fmt.Println("Pagamento aprovado!\n")
		return true
	}

	fmt.Println("Pagamento falhou!\n")
	return false
}

func menu() {
	cinema := carregar()

	for {
		fmt.Println("=== CINEMA (GO) ===")
		fmt.Println("1 - Ver filmes")
		fmt.Println("2 - Comprar ingresso")
		fmt.Println("0 - Sair")

		var op int
		fmt.Print("Escolha: ")
		fmt.Scanln(&op)

		switch op {
		case 1:
			listarFilmes(cinema)

		case 2:
			listarFilmes(cinema)

			var escolha int
			fmt.Print("Escolha o filme: ")
			fmt.Scanln(&escolha)

			if escolha <= 0 || escolha > len(cinema.Filmes) {
				fmt.Println("Opção inválida!\n")
				continue
			}

			filme := &cinema.Filmes[escolha-1]

			if escolherAssento(filme) {
				if pagamento() {
					salvar(cinema)
					fmt.Println("Compra finalizada!\n")
				}
			}

		case 0:
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida!\n")
		}
	}
}

func main() {
	menu()
}
