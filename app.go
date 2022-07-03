package main

import (
	"app/services"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	// "github.com/aws/jsii-runtime-go"
)

func main() {
	app := awscdk.NewApp(nil)

	services.TestRootStack(app, "AppStack")

	app.Synth(nil)
}
