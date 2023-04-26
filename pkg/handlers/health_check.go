package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HealthCheckHandler struct {}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method != "GET" {
		message := map[string]string{
			"message": "Method not allowed!",
		}

		jsonMessage, err := json.Marshal(message)

		if err != nil {
			fmt.Println("Error:", err)
		}
		
		w.WriteHeader(http.StatusMethodNotAllowed)

		w.Write([]byte(jsonMessage))

		return
	}

	message := map[string]string{
		"status": "Alive",
	}

	jsonMessage, err:= json.Marshal(message)

	if err != nil {
		log.Panic("Erro ao converter mensagem para JSON: ", err)
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonMessage))
}