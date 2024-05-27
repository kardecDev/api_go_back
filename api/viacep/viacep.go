package viacep

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kardecDev/api_go_back/config/env"
)

type ViaCepResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAF        string `json:"siaf"`
}

func GetCep(cep string) (*ViaCepResponse, error) {

	url := fmt.Sprintf("%s/%s/json", env.Env.ViaCepURL, cep)

	var ViaCepResponse ViaCepResponse

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ViaCepResponse)
	if err != nil {
		return nil, err
	}
	if ViaCepResponse.CEP == "" {
		return nil, fmt.Errorf("cep not found")
	}

	return &ViaCepResponse, nil

}
