package routers

import (
	"errors"
	"strings"

	"github.com/DainerAcosta/twittor/bd"
	"github.com/DainerAcosta/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string

var IDUsuario string

/*ProcesoToken procesa el token*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("este_es_mi_secreto")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
