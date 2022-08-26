package repository

import (
	"github.com/antony00/database"
	"github.com/antony00/model"
)

func InsertUsuario(usuario model.Usuario) error {
	db := database.GrabDB()

	statement := db.MustBegin()

	statement.MustExec("INSERT INTO usuario VALUES($1,$2)", usuario.Nome, usuario.RA)
	if err := statement.Commit(); err != nil {
		return err
	}

	return nil
}
