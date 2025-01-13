package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type BrasilApi struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type ApiEndereco struct {
	api      string
	endereco string
}

func main() {
	channel := make(chan ApiEndereco)

	go func() {
		channel <- ApiEndereco{"viaCep", viaCep()}
	}()

	go func() {
		channel <- ApiEndereco{"brasilApi", brasilApi()}
	}()

	select {
	case msg := <-channel:
		fmt.Printf("Recebido de %s o endereço:\n  %s\n", msg.api, msg.endereco)
	case <-time.After(time.Second * 1):
		panic("Timeout - Nenhuma API respondeu dentro de 1 segundo")
	}

}

func viaCep() string {
	// time.Sleep(time.Second)
	req, err := http.Get("http://viacep.com.br/ws/01153000/json/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v \n", err)
	}
	defer req.Body.Close()

	response, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v \n", err)
	}
	var data ViaCEP
	err = json.Unmarshal(response, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer parce da resposta: %v \n", err)
	}
	return fmt.Sprint(data)

}

func brasilApi() string {
	// time.Sleep(time.Second)
	req, err := http.Get("https://brasilapi.com.br/api/cep/v1/01153000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v \n", err)
	}
	defer req.Body.Close()

	response, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v \n", err)
	}
	var data BrasilApi
	err = json.Unmarshal(response, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v \n", err)
	}
	return fmt.Sprint(data)

}
