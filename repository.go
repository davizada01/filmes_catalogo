package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Filme struct {
	ID 			  string   `json:"id"`
	Titulo        string   `json:"titulo"`
	Ano           int      `json:"ano"`
	Duracao       int	   `json:"duracao"`
	Generos       []string `json:"genero"`
	Diretor       string  		`json:"diretor"`
	Edicao        string   `json:"edicao"`
	Formato       string   `json:"formato"`
	Idioma        string   `json:"idioma"`
	Classificacao string `json:"classificacao"`
}

var catalogo []Filme
const arquivoBD = "meus_filmes.json"

func carregarDados() {
	dados, err := os.ReadFile(arquivoBD)
	if err != nil {
		fmt.Println("AVISO: O Arquivo JSON não foi encontrado: ", err)
		fmt.Print("Verifique se o terminal está na pasta correta.\n\n")
		return
	}
	json.Unmarshal(dados, &catalogo)

	teveAlteracao := false
	for i := range catalogo {
		if catalogo[i].ID == "" {
			catalogo [i].ID = gerarID()
			teveAlteracao = true
		}
	}
	if teveAlteracao {
		salvarDados()
	}
}

func salvarDados() {
	dados, err := json.MarshalIndent(catalogo, "", " ")
	if err != nil {
		fmt.Println("Erro ao converter os dados: ", err)
		return
	}
	os.WriteFile(arquivoBD, dados, 0644)
}

func gerarID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}