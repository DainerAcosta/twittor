package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DainerAcosta/twittor/bd"
	"github.com/DainerAcosta/twittor/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "El password debe tener al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Usuario existente", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Error inesperado"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Error inesperado"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
