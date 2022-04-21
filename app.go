package main

import (
	"app/env"
	"app/stacks"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	// "github.com/aws/jsii-runtime-go"
)

func main() {
	app := awscdk.NewApp(nil)

	stacks.EksStack(app, "AppStack", (*stacks.AppStackProps)(env.Props))

	app.Synth(nil)
}
