package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

// This function initialize an AWS configuration using a region as a parameter
func StartAWS(region string) {
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion(region))

	if err != nil {
		panic("Error loading config: " + err.Error())
	}
}
