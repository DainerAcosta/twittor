package main

import (
	"log"

	"github.com/DainerAcosta/twittor/bd"
	"github.com/DainerAcosta/twittor/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
