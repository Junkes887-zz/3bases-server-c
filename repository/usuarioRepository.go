package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Junkes887/3bases-server-c/model"
	"github.com/olivere/elastic/v7"
)

type Client struct {
	DB  *elastic.Client
	CTX context.Context
}

func (client Client) FindAll() []model.Usuario {
	var usuarios []model.Usuario
	res, err := client.DB.Search().Index("usuarios").Do(client.CTX)

	if err != nil {
		fmt.Println(err)
	}
	if res != nil {
		for _, hit := range res.Hits.Hits {
			var usuario model.Usuario
			err := json.Unmarshal(hit.Source, &usuario)
			usuario.ID = hit.Id
			if err != nil {
				fmt.Println(err)
			}

			usuarios = append(usuarios, usuario)
		}
	}

	return usuarios
}

func (client Client) Find(cpf string) model.Usuario {
	var usuario model.Usuario
	ctx := context.Background()

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("cpf", cpf))

	res, err := client.DB.Search().Index("usuarios").SearchSource(searchSource).Do(ctx)

	if err != nil {
		fmt.Println(err)
	}

	if res != nil {
		for _, hit := range res.Hits.Hits {
			err := json.Unmarshal(hit.Source, &usuario)

			usuario.ID = hit.Id

			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return usuario
}

func (client Client) Save(usuario model.Usuario) string {
	ctx := context.Background()
	dataJSON, err := json.Marshal(usuario)
	js := string(dataJSON)

	ind, err := client.DB.Index().
		Index("usuarios").
		BodyJson(js).
		Type("usuario").
		Do(ctx)

	if err != nil {
		panic(err)
	}

	return ind.Id
}

func (client Client) Upadate(id string, usuario model.Usuario) string {
	res, err := client.DB.Update().Index("usuarios").Type("usuario").Id(id).
		Doc(map[string]interface{}{
			"cpf":                    usuario.CPF,
			"ultimaConsulta":         usuario.UltimaConsulta,
			"movimentacaoFinanceira": usuario.MovimentacaoFinanceira,
			"listDadosUltimaCompra":  usuario.ListDadosUltimaCompra,
		}).Do(client.CTX)

	if err != nil {
		fmt.Println(err)
	}

	if res == nil {
		return "Usuario não encontrado"
	}

	return "Usuario atualizado"
}

func (cliente Client) Delete(id string) string {
	res, err := cliente.DB.Delete().Index("usuarios").Type("usuario").Id(id).Do(cliente.CTX)

	if err != nil {
		fmt.Println(err)
	}

	if res == nil {
		return "Usuario não encontrado"
	}

	return "Usuario removido"
}
