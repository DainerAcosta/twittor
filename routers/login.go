package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DainerAcosta/twittor/bd"
	"github.com/DainerAcosta/twittor/jwt"
	"github.com/DainerAcosta/twittor/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o clave invalidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email requerido", 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "Usuario y/o clave invalidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Error al intentar generar token "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
