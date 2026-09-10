package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func adicionarFilme(scanner *bufio.Scanner) {
	var novoFilme Filme

   for {	
	fmt.Println("Título: ")
	scanner.Scan()
	titulo := strings.TrimSpace(scanner.Text())

	if titulo != "" {
		novoFilme.Titulo = titulo
		break
	}
	fmt.Println("O título não pode ficar em branco. Tente de novo") 
}
   
   for {
	fmt.Println("Ano de Lançamento: ")
	scanner.Scan()
	anoTexto := scanner.Text()
	anoNumero, err := strconv.Atoi(strings.TrimSpace(anoTexto))
	if err == nil && anoNumero > 1880 && anoNumero <= time.Now().Year()+5 {
	novoFilme.Ano = anoNumero
	break
	}
	fmt.Print("Digite um ano válido.\n")
   }
   for {
	fmt.Println("Tempo de duração total em MINUTOS (Ex.: 130): ")
	scanner.Scan()
	duracaoTexto := scanner.Text()
	duracaoNumero, err := strconv.Atoi(strings.TrimSpace(duracaoTexto))

	if err == nil && duracaoNumero > 0 {
		novoFilme.Duracao = duracaoNumero
		break
	}
	fmt.Println("Erro: Apenas números (em minutos)")
}

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

	novoFilme.ID = gerarID()

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
		fmt.Printf("%d. %s - %d | Duração: %dmin | (Gêneros: %s) | (Diretor: %s) | Mídia: %s - %s - %s | %s\n", i+1, filme.Titulo, filme.Ano, filme.Duracao, generosFormados, filme.Diretor, filme.Formato, filme.Edicao, filme.Idioma, filme.Classificacao)

	}
}

func removerFilme(scanner *bufio.Scanner) {
	if len(catalogo) == 0 {
		fmt.Println("O catálogo está vazio!")
		return
	}

	listarFilme()

	fmt.Print("\nDigite o número do filme que quer excluir (0 para voltar): ")
	scanner.Scan()
	input := scanner.Text()

	numeroDigitado, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil || numeroDigitado < 0 || numeroDigitado > len(catalogo) {
		fmt.Println("Número inválido. Selecione um número válido da lista.")
		return
	}
	if numeroDigitado == 0 {
		fmt.Println("\nExclusão foi cancelada")
		return
	}

	indice := numeroDigitado - 1
	idParaRemover := catalogo[indice].ID

	fmt.Printf("Filme '%s' foi excluído da lista.\n\n", catalogo[indice].Titulo)

	var novoCatalogo []Filme
	for _, filme := range catalogo {
		if filme.ID != idParaRemover {
			novoCatalogo = append(novoCatalogo, filme)
		}
	}
	
	catalogo = novoCatalogo
	salvarDados()
}

func buscarFilme(scanner *bufio.Scanner) {
	if len(catalogo) == 0 {
		fmt.Println("\nO catálogo está vazio!")
		return
	}

	fmt.Println("\nDigite um termo de busca (filme, diretor, formato ou gênero): ")
	scanner.Scan()
	textoBusca := scanner.Text()

	termoBuscaLimpo := removerAcentos(strings.ToLower(textoBusca))
	encontrou := false
	fmt.Println("----Resultados da Busca----")

	for _, filme := range catalogo {
		tituloLimpo := removerAcentos(strings.ToLower(filme.Titulo))
		diretorLimpo := removerAcentos(strings.ToLower(filme.Diretor))
		formatoLimpo := removerAcentos(strings.ToLower(filme.Formato))

		encontrouNoGenero := false
		for _, genero := range filme.Generos {
			generoLimpo := removerAcentos(strings.ToLower(genero))
			if strings.Contains(generoLimpo, termoBuscaLimpo) {
				encontrouNoGenero = true
			}
		}

		if strings.Contains(tituloLimpo, termoBuscaLimpo) || strings.Contains(diretorLimpo, termoBuscaLimpo) || strings.Contains(formatoLimpo, termoBuscaLimpo) || encontrouNoGenero {
			fmt.Println("Nome: ", filme.Titulo)
			fmt.Println("Diretor: ", filme.Diretor)
			fmt.Println("Gêneros: ", strings.Join(filme.Generos, ", "))
			fmt.Println("Formato: ", filme.Formato)
			fmt.Println("-------------------------")
			encontrou = true
		}
	}

	if !encontrou{
		fmt.Println("Nenhum filme encontrado!!")
		fmt.Println("Digite nome do FILME, DIRETOR, FORMATO OU GÊNERO que você quer encontrar!")
	}
}
	
	func removerAcentos (texto string) string {
		t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

		textoSemAcento, _, _ := transform.String(t, texto)

		return textoSemAcento
}
