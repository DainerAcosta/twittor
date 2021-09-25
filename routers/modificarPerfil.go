package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DainerAcosta/twittor/bd"
	"github.com/DainerAcosta/twittor/models"
)

/*ModificarPerfil modifica el perfil del usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	_, _, IDUsuario, _ := ProcesoToken(r.Header.Get("Authorization"))
	fmt.Println(IDUsuario)

	var status bool

	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "error inesperado"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se logro modificar el usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
