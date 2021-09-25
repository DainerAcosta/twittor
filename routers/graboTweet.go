package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DainerAcosta/twittor/bd"
	"github.com/DainerAcosta/twittor/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "error al intentar insertar el registro"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se logro insertar el tweet", 400)
	}

	w.WriteHeader(http.StatusCreated)

}
