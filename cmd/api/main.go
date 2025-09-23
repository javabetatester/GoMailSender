package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	//PARAMETRO POR QUERY
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		product := r.URL.Query().Get("product")
		id := r.URL.Query().Get("id")
		if product != "" || id != "" {
			w.Write([]byte(product + " " + id))
		}
	})

	//PARAMETRO POR PATH
	r.Get("/{viadoName}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "viadoName")
		w.Write([]byte(param + " da o cu"))
	})

	//PARAMETRO PARA JSON
	r.Get("/json/{ehViado}/{comeTraveco}", func(w http.ResponseWriter, r *http.Request) {

		// SETA QUE O HEADER VAI SER ...
		w.Header().Set("Content-Type", "application/json")

		// PEGA O PARAMETRO DA URL
		ehViado := chi.URLParam(r, "ehViado")
		comeTraveco := chi.URLParam(r, "comeTraveco")

		// CRIA UM MAP DE STRINGS, PARA REPRESENTAR CHAVE VALOR DA BODY DA RESPONSE
		obj := map[string]string{"message": "sucess", ehViado: "viado", comeTraveco: "come traveco"}
		// CONVERTE O MAP DE STRINGS CRIADO EM UM JSON
		b, _ := json.Marshal(obj)

		// RETORNA OS VALORES SETADOS NO OBJETO
		w.Write(b)
	})

	r.Get("/jsonChi/{ehViado}/{comeTraveco}", func(w http.ResponseWriter, r *http.Request) {

		// O CHI RENDER DESCARTA O USO DE CONTENT TYPE
		//w.Header().Set("Content-Type", "application/json")

		// PEGA O PARAMETRO DA URL
		ehViado := chi.URLParam(r, "ehViado")
		comeTraveco := chi.URLParam(r, "comeTraveco")

		// CRIA UM MAP DE STRINGS, PARA REPRESENTAR CHAVE VALOR DA BODY DA RESPONSE
		obj := map[string]string{"message": "sucess", ehViado: "viado", comeTraveco: "come traveco"}

		// RENDERIZA COMO RESPONSE
		render.JSON(w, r, obj)
		jsonData, _ := json.Marshal(obj)
		fmt.Printf("JSON enviado: %s\n", obj)      // JEITO PORCO
		fmt.Printf("JSON enviado: %s\n", jsonData) // JEITO ESBELTO
	})

	type travequeiros struct {
		CPF           string
		Nome          string
		NmrDeTravecos int
	}

	listaTravecos := make(map[string]travequeiros)

	r.Post("/perfilTravequeiro", func(w http.ResponseWriter, r *http.Request) {
		var travequeiro travequeiros
		render.DecodeJSON(r.Body, &travequeiro)
		listaTravecos[travequeiro.CPF] = travequeiro
		render.JSON(w, r, travequeiro)
	})

	r.Patch("/perfilTravequeiro/{cpf}", func(w http.ResponseWriter, r *http.Request) {
		cpf := chi.URLParam(r, "cpf")
		var travequeiro travequeiros
		render.DecodeJSON(r.Body, &travequeiro)
		for _, value := range listaTravecos {
			if value.CPF == cpf {
				listaTravecos[cpf] = travequeiro
			}
		}
		for _, value := range listaTravecos {
			fmt.Println(value.CPF)
		}

	})

	http.ListenAndServe(":4889", r)
}
