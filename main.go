package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/berpeda/userRegistration/awsgo"
	"github.com/berpeda/userRegistration/database"
	"github.com/berpeda/userRegistration/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	const region = "eu-west-1"

	// Initialize AWS in the specific region
	awsgo.StartAWS(region)

	// Verifies if the parameters are valid
	if !ValidParameters() {
		fmt.Println("Error at the parameters. It should send 'SecretName'")
		err := errors.New("error at the parameters, it should send SecretName")
		return event, err
	}

	// Gets the user's attributes
	userAttributes := event.Request.UserAttributes
	var datos models.SignUp
	datos.UserUUID = userAttributes["sub"]
	datos.UserEmail = userAttributes["email"]

	// Prints the user's information
	fmt.Println("The assigned UUID -> ", datos.UserUUID)
	fmt.Println("The assigned email -> ", datos.UserEmail)
	fmt.Println("The user email is verified -> ", userAttributes["email_verified"])

	// Reads the secret of the database
	// If there is an error, it will print the error
	err := database.ReadScecret()
	if err != nil {
		fmt.Println("Error reading the Secret " + err.Error())
	}

	err = database.SignUp(datos)

	if err != nil {
		fmt.Println("Error in the Sign Up")
	} else {
		fmt.Println("Sign Up succesfull!")
	}

	return event, err
}

// This functions verifies if the environment variable 'SecretName' is established
func ValidParameters() bool {
	_, getParam := os.LookupEnv("SecretName")
	return getParam
}
