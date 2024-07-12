package database

import (
	"fmt"

	"github.com/berpeda/userRegistration/models"
	"github.com/berpeda/userRegistration/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(user models.SignUp) error {
	fmt.Println("Registering a user...")

	err := DatabaseConnect()
	if err != nil {
		return err
	}

	defer Database.Close()

	_, err = Database.Exec(`INSERT INTO Usuarios (UUID_usuario, Email, Creado) VALUES (?, ?, ?)`, user.UserUUID, user.UserEmail, tools.DateMySQL())

	if err != nil {
		return err
	}

	return nil
}
