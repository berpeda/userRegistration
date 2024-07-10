package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/berpeda/userRegistration/models"
	secretsmanager "github.com/berpeda/userRegistration/secretsManager"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Database *sql.DB

func ReadScecret() error {
	SecretModel, err = secretsmanager.GetSecrets(os.Getenv("SecretName"))
	return err
}

func DatabaseConnect() error {
	Database, err = sql.Open("mysql", ConnectionString(SecretModel))
	if err != nil {
		fmt.Println("The connection to the database failed -> ", err.Error())
		return err
	}

	err = Database.Ping()
	if err != nil {
		fmt.Println("Ping got an error -> ", err)
	}

	fmt.Println("The connection is succesfull!")
	return nil
}

func ConnectionString(keys models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = keys.Username
	authToken = keys.Password
	dbEndpoint = keys.Host
	dbName = "comercialbermudez"
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)

}
