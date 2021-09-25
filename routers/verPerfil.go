package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DainerAcosta/twittor/bd"
)

func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID requerido", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Error inesperado"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
