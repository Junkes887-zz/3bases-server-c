package model

import "time"

type DadosUltimaCompra struct {
	Dado string `json:"dado"`
}

type Usuario struct {
	ID                     string              `json:"id"`
	CPF                    string              `json:"cpf"`
	UltimaConsulta         time.Time           `json:"ultimaConsulta"`
	MovimentacaoFinanceira float64             `json:"movimentacaoFinanceira"`
	ListDadosUltimaCompra  []DadosUltimaCompra `json:"listDadosUltimaCompra"`
}
