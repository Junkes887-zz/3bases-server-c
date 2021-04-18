package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Junkes887/3bases-server-c/model"
	"github.com/Junkes887/3bases-server-c/repository"
	"github.com/julienschmidt/httprouter"
	"github.com/olivere/elastic/v7"
)

type Client struct {
	DB  *elastic.Client
	REP repository.Client
}

func (client Client) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	usuarios := client.REP.FindAll()

	js, err := json.Marshal(usuarios)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (client Client) Find(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cpf := p.ByName("cpf")

	usuario := client.REP.Find(cpf)

	js, err := json.Marshal(usuario)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (client Client) Save(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var usuario model.Usuario

	json.NewDecoder(r.Body).Decode(&usuario)
	create := client.REP.Save(usuario)

	js, err := json.Marshal(create)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(js)
}

func (client Client) Upadate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var usuario model.Usuario
	id := p.ByName("id")

	json.NewDecoder(r.Body).Decode(&usuario)

	menssage := client.REP.Upadate(id, usuario)

	js, err := json.Marshal(menssage)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (cliente Client) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	menssage := cliente.REP.Delete(id)

	js, err := json.Marshal(menssage)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
