package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	carregarDados()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n---|Catálogo de filmes mídia física|---")
		fmt.Println("\n1. Adicionar um novo filme")
		fmt.Println("2. Listar filmes")
		fmt.Println("3. Excluir filme")
		fmt.Println("4. Buscar Filme")
		fmt.Println("5. Sair")
		fmt.Print("\nEscolha uma opção:")

		scanner.Scan()
		opcao := scanner.Text()
		switch strings.TrimSpace(opcao) {

		case "1":
			adicionarFilme(scanner)
		case "2":
			listarFilme()
		case "3":
			removerFilme(scanner)
		case "4":
			buscarFilme(scanner)
		case "5":
			fmt.Println("\nSaindo do catálogo...")
			return
		default:
			fmt.Println("Opção Inválida, tenta de novo")
		}

	}
}

