package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"projetoBuscaCepCli/models"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Busca cep usando uma api gratuita da internet"
	app.Usage = "Busca cep e tràs dados para usuário pelo terminal"

	flag := []cli.Flag{
		cli.StringFlag{
			Name:  "cep",
			Value: "41900205",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "buscar",
			Usage:  "Buscar endereco da internet a paritr de cep",
			Flags:  flag,
			Action: buscarCep,
		},
	}

	return app
}

func buscarCep(c *cli.Context) {
	cep := c.String("cep")
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	req, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro ao fazer a requisição: %v\n", err)
		return
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Erro ao ler a resposta da requisição: %v\n", err)
		return
	}

	var viaCep models.Endereco
	err = json.Unmarshal(res, &viaCep)
	if err != nil {
		fmt.Printf("Erro ao fazer o parse da resposta da requisição: %v\n", err)
		return
	}

	fmt.Println(viaCep)
}
