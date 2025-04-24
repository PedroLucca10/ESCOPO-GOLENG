package main

import "fmt"

var nomeEscola = "Escola técnica SENAI"

func main(){
	nome := "lucas"
	idade := 16

    mensagem := boasVindas(nome)
    fmt.Println(mensagem)

	status := verificaMaioridade(idade)
	fmt.Println(status)
}