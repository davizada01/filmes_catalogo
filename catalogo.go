package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Filme struct {
	Titulo        string   `json:"titulo"`
	Ano           int      `json:"ano"`
	Duracao       string   `json:"duracao"`
	Generos       []string `json:"genero"`
	Diretor       string   `json:"diretor"`
	Edicao        string   `json:"edicao"`
	Formato       string   `json:"formato"`
	Idioma        string   `json:"idioma"`
	Classificacao string `json:"classificacao"`
}

var catalogo []Filme

const arquivoBD = "meus_filmes.json"

func main() {
	carregarDados()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n---|Catálogo de filmes mídia física|---")
		fmt.Println("\n1. Adicionar um novo filme")
		fmt.Println("2. Listar filmes")
		fmt.Println("3. Excluir filme")
		fmt.Println("4. Sair")
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
			fmt.Println("\nSaindo do catálogo...")
			return
		default:
			fmt.Println("Opção Inválida, tenta de novo")
		}

	}

}

func carregarDados() {
	dados, err := os.ReadFile(arquivoBD)
	if err != nil {
		return
	}
	json.Unmarshal(dados, &catalogo)
}

func salvarDados() {
	dados, err := json.MarshalIndent(catalogo, "", " ")
	if err != nil {
		fmt.Println("Erro ao converter os dados: ", err)
		return
	}
	os.WriteFile(arquivoBD, dados, 0644)
}

func adicionarFilme(scanner *bufio.Scanner) {
	var novoFilme Filme

	fmt.Println("Título: ")
	scanner.Scan()
	novoFilme.Titulo = scanner.Text()

	fmt.Println("Ano de Lançamento: ")
	scanner.Scan()
	anoTexto := scanner.Text()
	anoNumero, err := strconv.Atoi(strings.TrimSpace(anoTexto))
	if err != nil {
		novoFilme.Ano = 0
	} else {
		novoFilme.Ano = anoNumero
	}

	fmt.Println("Tempo de duração (Xh XXmin): ")
	scanner.Scan()
	novoFilme.Duracao = scanner.Text()

	fmt.Println("Gêneros (separe com vírgula): ")
	scanner.Scan()
	generosTexto := scanner.Text()

	listaDeGeneros := strings.Split(generosTexto, ",")

	for i := range listaDeGeneros {
		listaDeGeneros[i] = strings.TrimSpace(listaDeGeneros[i])
	}
	novoFilme.Generos = listaDeGeneros

	fmt.Println("Diretor: ")
	scanner.Scan()
	novoFilme.Diretor = scanner.Text()

	fmt.Println("Edição: ")
	scanner.Scan()
	novoFilme.Edicao = scanner.Text()

	fmt.Println("Formato: ")
	scanner.Scan()
	novoFilme.Formato = scanner.Text()

	fmt.Println("Idioma: ")
	scanner.Scan()
	novoFilme.Idioma = scanner.Text()

	fmt.Println("Classificação Indicativa (Ex.: 14 (BR) ou PG-13 (USA)): ")
	scanner.Scan()
	novoFilme.Classificacao = scanner.Text()

	catalogo = append(catalogo, novoFilme)
	salvarDados()
	fmt.Println("\nFilme foi adicionado!")

}

func listarFilme() {
	if len(catalogo) == 0 {
		fmt.Println("\nO catálogo ainda está vazio.")
		return
	}
	fmt.Println("\n--Filmes--")
	for i, filme := range catalogo {
		generosFormados := strings.Join(filme.Generos, ", ")
		fmt.Printf("%d. %s - %d | Duração: %s | (Gêneros: %s) | (Diretor: %s) | Mídia: %s - %s - %s | %s\n", i+1, filme.Titulo, filme.Ano, filme.Duracao, generosFormados, filme.Diretor, filme.Formato, filme.Edicao, filme.Idioma, filme.Classificacao)

	}
}

func removerFilme(scanner *bufio.Scanner) {
	if len(catalogo) == 0 {
		fmt.Println("O catálogo está vazio! ")
		return
	}

	listarFilme()

	fmt.Print("\nDigite o número do filme que quer excluir (0 para voltar): ")
	scanner.Scan()
	input := scanner.Text()

	numeroDigitado, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil || numeroDigitado < 0 || numeroDigitado > len(catalogo) {
		fmt.Println("Número ainda não foi registrado. Selecione um número válido")
		return
	}
	if numeroDigitado == 0 {
		fmt.Println("\nExclusão foi cancelada")
		return
	}

	indice := numeroDigitado - 1

	fmt.Printf("Filme '%s' foi excluído da lista.\n\n", catalogo[indice].Titulo)

	catalogo = append(catalogo[:indice], catalogo[indice+1:]...)
	salvarDados()
}
