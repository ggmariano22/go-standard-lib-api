package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Product struct {
	Name  string `json:"name"`
	Price float32 `json:"price"`
	Type  string `json:"type"`
}

type CreateProductHandler struct {}

func NewCreateProductHandler() *CreateProductHandler {
	return &CreateProductHandler{}
}

func (c *CreateProductHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "Method not allowed."}`))
	}

	defer r.Body.Close()
	var content Product

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Erro ao ler body da request: ", err)
	}

	err = json.Unmarshal(body, &content)

	if err != nil {
		fmt.Println("Erro ao criar produto: ", err)
	}

	fmt.Println("Body recebido é:")
	fmt.Println(content)
}