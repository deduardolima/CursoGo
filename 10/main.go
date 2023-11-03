package main

import "fmt"

type Endereco struct {
	Logadouro string
	Numero int
	Cidade string
	Estado string
}

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Client) Desativar(){
		c.Ativo = false
		fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func main() {
	eduardo := Client{
		Nome:  "Eduardo",
		Idade: 30,
		Ativo: true,
	}
	eduardo.Idade = 31
	eduardo.Estado = "Paraná"
	eduardo.Cidade = "Curitiba"
	eduardo.Desativar()
}
