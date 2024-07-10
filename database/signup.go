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

	sentence := "INSERT INTO Usuarios (UUID_usuario, Email, Creado) VALUES ('" + user.UserUUID + "', '" + user.UserEmail + "', " + tools.DateMySQL() + ");"
	fmt.Println(sentence)

	_, err = Database.Exec(sentence)

	if err != nil {
		return err
	}

	return nil
}
